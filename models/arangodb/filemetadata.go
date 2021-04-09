//+build

package arangodb

import "time"

type FileMetadata struct {
	Id       string `json:"id"`
	FileId   string `json:"-"`
	BucketId string `json:"bucket_id"`
	Path     string `json:"path"`
	Name     string `json:"name"`

	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
	IsHidden    bool   `json:"is_hidden"`

	IsDeleted   bool      `json:"-"`
	DeletedDate time.Time `json:"-"`

	UploadedDate time.Time `json:"-"`
	ExpiredDate  time.Time `json:"expired_date"`
}

type fileMetadata struct {
	FileId   string `json:"fid"`
	BucketId string `json:"bucket_id"`
	Path     string `json:"path"`
	Name     string `json:"name"`

	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
	IsHidden    bool   `json:"is_hidden"`

	IsDeleted   bool      `json:"is_deleted"`
	DeletedDate time.Time `json:"deleted_date"`

	UploadedDate time.Time `json:"upload_date"`
	ExpiredDate  time.Time `json:"expired_date"`
}
