package models

import "time"

type Datafile struct {
	ID          string     `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	ContentType string     `json:"content_type" db:"content_type"`
	Size        int64      `json:"size" db:"size"`
	Path        string     `json:"path" db:"path"`
	URL         string     `json:"url" db:"url"`
	UploadedBy  int64      `json:"uploaded_by" db:"uploaded_by"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
