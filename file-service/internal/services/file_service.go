package services

import (
	"context"
	"file-service/internal/models"
	"file-service/internal/repos"
	"file-service/internal/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const maxFileSizeBytes = 500 << 20 // 500MB

type FileService struct {
	store *storage.LocalStorage
	repo  *repos.FileRepo
}

func NewFileService(store *storage.LocalStorage, repo *repos.FileRepo) *FileService {
	return &FileService{
		store: store,
		repo:  repo,
	}
}

func (s *FileService) Upload(ctx context.Context, name, ctype string, uploadedBy int64, content []byte) (*models.Datafile, error) {
	if int64(len(content)) > maxFileSizeBytes {
		return nil, fmt.Errorf("file size exceeds max file size %dMB", maxFileSizeBytes>>20)
	}

	free, err := s.store.FreeSpaceGB()
	if err == nil && free < 5.0 {
		return nil, fmt.Errorf("disk space is less than %.1fGB", free)
	}

	id := uuid.NewString()
	key := fmt.Sprintf("%d/%s/%s", uploadedBy, id, name)
	if err := s.store.Save(key, content); err != nil {
		return nil, fmt.Errorf("failed to write file to disk: %w", err)
	}

	file := &models.Datafile{
		ID:          id,
		Name:        name,
		ContentType: ctype,
		Size:        int64(len(content)),
		Path:        s.store.FullPath(key),
		URL:         s.store.PublicURL(key),
		UploadedBy:  uploadedBy,
		CreatedAt:   time.Now(),
	}

	_, err = s.repo.Save(ctx, file)
	if err != nil {
		s.store.Delete(key)
		return nil, fmt.Errorf("can't write file to db: %w", err)
	}

	return file, nil
}

func (s *FileService) GetFile(ctx context.Context, fileID string) (*models.Datafile, error) {
	return s.repo.FindByID(ctx, fileID)
}

func (s *FileService) GetContent(ctx context.Context, fileID string) (*models.Datafile, []byte, error) {
	file, err := s.repo.FindByID(ctx, fileID)
	if err != nil {
		return nil, nil, err
	}

	f, _, err := s.store.Open(s.storageKey(file))
	if err != nil {
		return nil, nil, fmt.Errorf("file not found: %w", err)
	}
	defer f.Close()

	content := make([]byte, file.Size)
	if _, err := f.Read(content); err != nil {
		return nil, nil, err
	}

	return file, content, nil
}

func (s *FileService) Delete(ctx context.Context, fileID string) error {
	file, err := s.repo.FindByID(ctx, fileID)
	if err != nil {
		return err
	}

	if err := s.store.Delete(s.storageKey(file)); err != nil {
		return err
	}

	return s.repo.Delete(ctx, fileID)
}

func (s *FileService) storageKey(f *models.Datafile) string {
	return fmt.Sprintf("%d/%s/%s", f.UploadedBy, f.ID, f.Name)
}
