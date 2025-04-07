package handlers

import (
	"ToGoList/internal/models"
	"ToGoList/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	db := c.MustGet("db").(*database.DB)
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	db := c.MustGet("db").(*database.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Invalid book's id",
		})
		return
	}
	var book models.Book
	db.Find(&book, "id = ?", id)
	c.JSON(http.StatusOK, book)
}

func AddBook(c *gin.Context) {
	db := c.MustGet("db").(*database.DB) // инструмент для работы Бд - тут ОРМ
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*database.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Couldn't parse an ID from url",
		})
		return
	}
	res := db.Where("id = ?", id).Delete(&models.Book{})
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": res.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book with id = " + strconv.Itoa(id) + " was deleted",
	})
}

func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*database.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Couldn't parse an ID from url",
		})
		return
	}
	var book models.Book
	db.First(&book, "id = ?", id)
	if book.Title == "" { // == nil
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't find book with this ID in library",
		})
	}

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Incorrect new data for book",
		})
	}
	db.Model(&book).Updates(updatedBook)
}
