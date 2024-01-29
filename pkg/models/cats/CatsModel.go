package models

import "gorm.io/gorm"

type Cats struct {
	gorm.Model
	Name   string `json:"name"`
	Race   string `json:"race"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}
