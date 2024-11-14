package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    ID                uint           `gorm:"primaryKey"`
    CreatedAt         time.Time
    UpdatedAt         time.Time
    DeletedAt         gorm.DeletedAt `gorm:"index"`
    MobileNumber      string         `gorm:"unique;not null"`
    DeviceFingerprint string
}