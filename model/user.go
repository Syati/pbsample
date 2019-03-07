package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
	Id     int64
	Name   string
}
