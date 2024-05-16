package routes

import (
	"books_go/handler"
	"books_go/server"
)

func Routing() {
	server.E.GET("/", handler.MainPage)
	server.E.GET("/book", handler.AllBooks)
	server.E.GET("/book/:id", handler.GetBook)
	server.E.POST("/book", handler.SaveBook)
	server.E.PUT("/book/:id", handler.UpdateBook)
	server.E.DELETE("/book/:id", handler.DeleteBook)
}
