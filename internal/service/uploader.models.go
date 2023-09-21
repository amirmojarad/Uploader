package service

import (
	"mime/multipart"
	"time"
)

type File struct {
	Id        uint
	Title     string
	Volume    float64
	CreatedAt time.Time
	Format    string
}

type UploadFileRequest struct {
	File *multipart.File
}

type UploadFileResponse struct {
	File
}
