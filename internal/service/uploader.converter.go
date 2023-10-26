package service

func toSvcBucketEntity(entity BucketEntity) {

}

func toFileEntity(fileName, fileType string, userId, bucketId uint, size int8) FileEntity {
	return FileEntity{
		Name:     fileName,
		UserID:   userId,
		BucketID: bucketId,
		Size:     size,
		Type:     fileType,
	}
}
