package repository

import "gorm.io/gorm"

type BucketEntity struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);not null"`
	UserID uint   `gorm:"type:bigint;not null"`
}

func (BucketEntity) TableName() string {
	return "buckets"
}
