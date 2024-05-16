package handler

import (
	"books_go/db"
	"books_go/models"
	"books_go/server"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func MainPage(c echo.Context) error {
	response := "Welcome. Available routes\n"
	routes := server.E.Routes()
	for _, route := range routes {
		response += fmt.Sprintf("%+v\n", *route)
	}
	return c.String(http.StatusOK, response)
}

func AllBooks(c echo.Context) error {
	books := []models.Book{}
	db.DB.Find(&books)
	fmt.Println(books)
	return c.JSON(http.StatusOK, &books)
}

func GetBook(c echo.Context) error {
	book := models.Book{}
	result := db.DB.First(&book, "id=?", c.Param("id"))
	if result.Error != nil {
		return c.String(http.StatusNotFound, result.Error.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("%+v", book))
}

func SaveBook(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	authorName := c.QueryParam("authorName")
	price, _ := strconv.Atoi(c.QueryParam("price"))
	result := db.DB.Create(&models.Book{ID: id, Name: name, AuthorName: authorName, Price: uint(price)})
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, result.Error.Error())
	}
	return c.String(http.StatusCreated, fmt.Sprintf("book of id:%v added", id))
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	book := models.Book{
		ID: id,
	}
	db.DB.First(&book)
	jsonMap := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	for key, value := range jsonMap {
		switch key {
		case "authorName":
			book.AuthorName = fmt.Sprint(value)
		case "name":
			book.Name = fmt.Sprint(value)
		case "price":
			price, _ := strconv.Atoi(fmt.Sprint(value))
			book.Price = uint(price)
		}
	}
	result := db.DB.Save(&book)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, result.Error.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("updated entry: %+v", book))
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	book := models.Book{ID: id}
	result := db.DB.Delete(&book)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, result.Error.Error())
	}
	return c.String(http.StatusOK, "Entry with id: "+id+"deleted successfully")
}
