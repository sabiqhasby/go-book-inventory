package main

import (
	"book-inventory/app"
	"book-inventory/conn"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := conn.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	handler := app.New(conn)
	r.GET("/books", handler.GetBooks)

	r.Run(":8080")

}
