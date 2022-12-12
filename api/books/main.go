package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID     string `json:id`
	Name   string `json:name`
	Author string `json:author`
	Link   string `json:link`
}

var books = []book{
	{
		ID:     "0",
		Name:   "SCRUM",
		Author: "Author SCRUM",
		Link:   "google.com",
	},

	{
		ID:     "1",
		Name:   "Livro 2",
		Author: "Author 2",
		Link:   "google.com/2",
	},
}

func main() {
	router := gin.Default()
	router.GET("/books", getAllBooks)
	router.POST("/books", postBook)
	router.GET("/books/:id", getBookByID)
	router.DELETE("/books/:id", deleteBook)
	router.Run("localhost:8080")
}

func getAllBooks(c *gin.Context) {
	//MARK: Passando o status code e o objeto que será exibido em forma de JSON
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	//MARK: Esse parametro é passado na hora que registramos a rota, ou seja, não é no body e sim um /0, /1
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func postBook(c *gin.Context) {
	var newBook book
	if err := c.Bind(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, books)
}

//MARK: Me desafiei a fazer esse endpoint sozinho :))
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	var newList = []book{}

	for _, b := range books {
		if b.ID != id {
			newList = append(newList, b)
		}
	}
	books = newList

	c.IndentedJSON(http.StatusAccepted, books)
}
