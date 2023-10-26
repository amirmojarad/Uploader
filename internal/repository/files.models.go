package repository

import "gorm.io/gorm"

type FileEntity struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	UserID   uint   `gorm:"type:bigint;not null"`
	BucketID uint   `gorm:"type:bigint;not null"`
	Size     int8   `gorm:"type:int8;not null"`
	Type     string `gorm:"type:varchar(255);not null"`
}

func (FileEntity) TableName() string {
	return "files"
}
