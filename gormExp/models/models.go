package models

import "gorm.io/gorm"

type Curso struct {
	gorm.Model
	Id        int
	Nome      string
	Area      string
	Professor string
}
