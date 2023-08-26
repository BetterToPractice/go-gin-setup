package models

import (
	"crypto/sha256"
	"encoding/hex"
	"gorm.io/gorm"
	"unsafe"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;size:64;not null;index;" json:"username" validate:"required"`
	Password string `gorm:"column:password;not null;" json:"password" validate:"required"`
	Email    string `gorm:"column:email;not null;" json:"email" validate:"required"`
}

func HashPassword(password string) string {
	sum := sha256.Sum256(*(*[]byte)(unsafe.Pointer(&password)))
	return hex.EncodeToString(sum[:])
}
