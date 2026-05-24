package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

type LocalStorage struct {
	path string
	url  string
}

func NewLocalStorage(path, url string) (*LocalStorage, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("can't create store folder: %w", err)
	}

	return &LocalStorage{
		path: path,
		url:  url,
	}, nil
}

func (s *LocalStorage) Save(key string, data []byte) error {
	fullPath := filepath.Join(s.path, key)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return err
	}
	return os.WriteFile(fullPath, data, 0644)
}

func (s *LocalStorage) SaveStream(key string, r io.Reader) (int64, error) {
	fullPath := filepath.Join(s.path, key)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return 0, err
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return io.Copy(f, r)
}

func (s *LocalStorage) Open(key string) (*os.File, int64, error) {
	fullPath := filepath.Join(s.path, key)
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, 0, err
	}
	info, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, 0, err
	}
	return f, info.Size(), nil
}

func (s *LocalStorage) Delete(key string) error {
	err := os.Remove(filepath.Join(s.path, key))
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func (s *LocalStorage) PublicURL(key string) string {
	return s.path + "/" + key
}

func (s *LocalStorage) FullPath(key string) string {
	return filepath.Join(s.path, key)
}

func (s *LocalStorage) FreeSpaceGB() (float64, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(s.path, &stat); err != nil {
		return 0, err
	}
	return float64(stat.Bavail*uint64(stat.Bsize)) / (1024 * 1024 * 1024), nil
}
