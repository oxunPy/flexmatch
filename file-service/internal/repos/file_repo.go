package repos

import (
	"context"
	"file-service/internal/database"
	"file-service/internal/models"
	"fmt"
	"time"
)

type FileRepo struct {
	storage *database.PostgresStorage
}

func NewFileRepo(storage *database.PostgresStorage) *FileRepo {
	return &FileRepo{
		storage: storage,
	}
}

func (r *FileRepo) Save(ctx context.Context, file *models.Datafile) (*models.Datafile, error) {
	var query = `
		INSERT INTO datafile (id, name, content_type, size, path, url, uploaded_by)
		VALUES ($1, $2, $3, $4, $5, $6, $6)
		RETURNING id, name, content_type, size, path, url, uploaded_by, created_at, deleted_at
	`

	var stored models.Datafile
	err := r.storage.
		QueryRow(
			ctx, query, file.ID, file.Name, file.ContentType, file.Size, file.Path, file.URL, file.UploadedBy,
		).
		Scan(
			&stored.ID, &stored.Name, &stored.ContentType, &stored.Size, &stored.Path, &stored.URL, &stored.UploadedBy, &stored.CreatedAt, &stored.DeletedAt,
		)
	if err != nil {
		return nil, err
	}

	return &stored, nil
}

func (r *FileRepo) FindByID(ctx context.Context, id string) (*models.Datafile, error) {
	var query = `
		SELECT id, name, content_type, size, path, url, uploaded_by, created_at, deleted_at
		FROM datafile
		WHERE id = $1
	`

	var file models.Datafile
	err := r.storage.
		QueryRow(ctx, query, id).
		Scan(&file.ID, &file.Name, &file.ContentType, &file.Size, &file.Path, &file.URL, &file.UploadedBy, &file.CreatedAt, &file.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *FileRepo) Delete(ctx context.Context, id string) error {
	var query = `
		UPDATE datafile 
		SET deleted_at = $1 
		WHERE id = $2 AND deleted_at IS NULL
	`
	ct, err := r.storage.Exec(ctx, query, time.Now(), id)
	if err != nil || ct.RowsAffected() == 0 {
		return fmt.Errorf("can't delete file at id: %s", id)
	}

	return nil
}
