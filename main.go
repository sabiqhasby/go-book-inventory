package main

import (
	"book-inventory/app"
	"book-inventory/auth"
	"book-inventory/conn"
	"book-inventory/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := conn.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	handler := app.New(conn)

	//HOME
	r.GET("/", auth.HomeHandler)

	//login
	r.GET("/login", auth.LoginGetHandler)
	r.POST("/login", auth.LoginPostHandler)

	//Get all books
	r.GET("/books", middlewares.AuthValid, handler.GetBooks)
	r.GET("/book/:id", middlewares.AuthValid, handler.GetBookById)
	//Insert Book
	r.GET("/add-book", middlewares.AuthValid, handler.AddBook)
	r.POST("/book", middlewares.AuthValid, handler.PostBook)

	//update book
	r.GET("/update-book/:id", middlewares.AuthValid, handler.UpdateBook)
	r.POST("/update-book/:id", middlewares.AuthValid, handler.PutBook)

	//Delete book
	r.POST("/delete-book/:id", middlewares.AuthValid, handler.DeleteBook)

	r.Run(":8080")

}
