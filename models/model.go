package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          int    `json:"id" form:"id" gorm:"primaryKey"`
	Title       string `json:"title" form:"title" binding:"required"`
	Author      string `json:"author" form:"author"  binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Stock       int    `json:"stock"  form:"stock" binding:"required"`
}

type Login struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

const (
	USER     = "admin"
	PASSWORD = "Pass1234"
	SECRET   = "secret"
)
