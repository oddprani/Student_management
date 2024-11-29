package database

import "gorm.io/gorm"

/// Student management system

type UserType string

const (
	Admin   UserType = "admin"
	Teacher UserType = "teacher"
	Student UserType = "student"
)

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null" json:"username"`
	Password string    `gorm:"not null" json:"-"`
	Type     UserType  `gorm:"not null" json:"type"`
	Sessions []Session `json:"-"`
}

type Session struct {
	gorm.Model
	Token  string `gorm:"unique;not null"`
	UserID uint
}
