package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Etag string `json:"etag"`
}
