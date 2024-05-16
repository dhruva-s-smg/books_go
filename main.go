package main

import (
	"books_go/db"
	"books_go/handler"
	"books_go/server"
	"os"

	"github.com/joho/godotenv"
)

//var E = echo.New()

func main() {
	godotenv.Load()
	// dsn := os.Getenv("DB_URL")
	// if dsn == "" {
	// 	log.Fatal("No DB_URL found in environvent")
	// }

	// DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// DB.AutoMigrate(&models.Book{})

	// mainPage := func(c echo.Context) error {
	// 	response := "Welcome. Available routes\n"
	// 	routes := E.Routes()
	// 	for _, route := range routes {
	// 		response += fmt.Sprintf("%+v\n", *route)
	// 	}
	// 	return c.String(http.StatusOK, response)
	// }

	// allBooks := func(c echo.Context) error {
	// 	books := []models.Book{}
	// 	DB.Find(&books)
	// 	fmt.Println(books)
	// 	return c.JSON(http.StatusOK, &books)
	// }

	// getBook := func(c echo.Context) error {
	// 	book := models.Book{}
	// 	result := DB.First(&book, "id=?", c.Param("id"))
	// 	if result.Error != nil {
	// 		return c.String(http.StatusNotFound, result.Error.Error())
	// 	}
	// 	return c.String(http.StatusOK, fmt.Sprintf("%+v", book))
	// }

	// saveBook := func(c echo.Context) error {
	// 	id := c.QueryParam("id")
	// 	name := c.QueryParam("name")
	// 	authorName := c.QueryParam("authorName")
	// 	price, _ := strconv.Atoi(c.QueryParam("price"))
	// 	result := DB.Create(&models.Book{ID: id, Name: name, AuthorName: authorName, Price: uint(price)})
	// 	if result.Error != nil {
	// 		return c.String(http.StatusInternalServerError, result.Error.Error())
	// 	}
	// 	return c.String(http.StatusCreated, fmt.Sprintf("book of id:%v added", id))
	// }

	// updateBook := func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	fmt.Println(id)
	// 	book := models.Book{
	// 		ID: id,
	// 	}
	// 	DB.First(&book)
	// 	jsonMap := make(map[string]interface{})
	// 	json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	// 	for key, value := range jsonMap {
	// 		switch key {
	// 		case "authorName":
	// 			book.AuthorName = fmt.Sprint(value)
	// 		case "name":
	// 			book.Name = fmt.Sprint(value)
	// 		case "price":
	// 			price, _ := strconv.Atoi(fmt.Sprint(value))
	// 			book.Price = uint(price)
	// 		}
	// 	}
	// 	result := DB.Save(&book)
	// 	if result.Error != nil {
	// 		return c.String(http.StatusInternalServerError, result.Error.Error())
	// 	}
	// 	return c.String(http.StatusOK, fmt.Sprintf("updated entry: %+v", book))
	// }

	// deleteBook := func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	book := models.Book{ID: id}
	// 	result := DB.Delete(&book)
	// 	if result.Error != nil {
	// 		return c.String(http.StatusInternalServerError, result.Error.Error())
	// 	}
	// 	return c.String(http.StatusOK, "Entry with id: "+id+"deleted successfully")
	// }

	// E.GET("/", mainPage)
	// E.GET("/book", allBooks)
	// E.GET("/book/:id", getBook)
	// E.POST("/book", saveBook)
	// E.PUT("/book/:id", updateBook)
	// E.DELETE("/book/:id", deleteBook)

	db.InitDB()
	server.E.GET("/", handler.MainPage)
	server.E.GET("/book", handler.AllBooks)
	server.E.GET("/book/:id", handler.GetBook)
	server.E.POST("/book", handler.SaveBook)
	server.E.PUT("/book/:id", handler.UpdateBook)
	server.E.DELETE("/book/:id", handler.DeleteBook)
	//fmt.Println("server port:", os.Getenv("PORT"))
	server.E.Logger.Fatal(server.E.Start(":" + os.Getenv("PORT")))
}
