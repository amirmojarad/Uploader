package service

import (
	"bytes"
	"mime/multipart"
	"time"
)

type UploadFileRequest struct {
	File     multipart.File
	FileName string
	UserId   uint
}

type UploadFileResponse struct {
	FileEntity
}

type AddFileRequest struct {
	FileEntity
}
type AddFileResponse struct {
	FileEntity
}

type GetAllFilesByIdRequest struct {
}

type GetAllFilesByIdResponse struct {
}

type GetFileByIdRequest struct {
	UserId uint
	FileId uint
}

type GetFileByIdResponse struct {
	File *bytes.Buffer
}

type AddBucketRequest struct {
	BucketName string
	UserId     uint
}

type AddBucketResponse struct {
	BucketEntity
}

type GetBucketByIdRequest struct {
}

type GetBucketByIdResponse struct {
	BucketEntity
}

type GetBucketByUserIdRequest struct {
	UserId uint
}

type GetBucketByUserIdResponse struct {
	BucketEntity
}

type GetFileByUserIdRequest struct {
	UserId uint
}

type GetFileByUserIdResponse struct {
	FileEntity
}

type DeleteBucketRequest struct {
	UserId uint
	Id     uint
}

type DeleteBucketResponse struct {
}

type Pagination struct {
	Index int
	Size  int
	Total int
}

type ListFilesRequest struct {
	UserId uint
	From   time.Time
	To     time.Time
	Pagination
}

type ListFilesResponse struct {
	Result []FileEntity
	Pagination
}

type FileEntity struct {
	Id       uint
	Name     string
	UserID   uint
	BucketID uint
	Size     int8
	Type     string
}

type BucketEntity struct {
	Id     uint
	Name   string
	UserID uint
}

type CreateBucketIfNotExistsRequest struct {
	BucketName string
	UserId     uint
}
type CreateBucketIfNotExistsResponse struct {
	BucketEntity
}
