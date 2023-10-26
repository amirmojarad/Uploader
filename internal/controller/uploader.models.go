package controller

import (
	"mime/multipart"
	"time"
)

type File struct {
	Id       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	UserID   uint   `json:"user_id,omitempty"`
	BucketID uint   `json:"bucket_id,omitempty"`
	Size     int8   `json:"size,omitempty"`
	Type     string `json:"type,omitempty"`
}

type UploadFileRequest struct {
	File     multipart.File
	FileName string
	UserId   uint
}

type UploadFileResponse struct {
	File
}

type GetFileRequest struct {
	Id uint `form:"id"`
}

type Pagination struct {
	Index int `url:"index"`
	Size  int `url:"size"`
	Total int `json:"total"`
}

type ListFilesRequest struct {
	UserId uint
	From   time.Time `url:"from"`
	To     time.Time `url:"to"`
	Pagination
}

type ListFilesResponse struct {
	Result []File `json:"result"`
	Pagination
}
