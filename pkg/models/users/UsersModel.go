package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
