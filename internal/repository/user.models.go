package repository

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Email          string `gorm:"type:varchar(255);not null;unique"`
	HashedPassword string `gorm:"type:varchar(255);not null"`
	RefreshToken   string `gorm:"type:varchar(255);not null"`
}

func (UserEntity) TableName() string {
	return "users"
}
