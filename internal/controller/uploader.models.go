package controller

import (
	"mime/multipart"
	"time"
)

type File struct {
	Id        uint      `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Volume    float64   `json:"volume,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Format    string    `json:"format,omitempty"`
}

type UploadFileRequest struct {
	File *multipart.File
}

type UploadFileResponse struct {
	File
}
