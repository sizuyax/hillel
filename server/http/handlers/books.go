package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hillel/errors"
	"hillel/logger"
	"net/http"
	"strconv"
	"strings"
)

var postRequestBook struct {
	Books           []string `json:"Books"`
	booksListString string
}

var putRequestBook struct {
	NewBooks           []string `json:"NewBooks"`
	newBooksListString string
}

func GetBooks(c echo.Context) error {

	if postRequestBook.Books == nil {
		return c.JSON(http.StatusBadRequest, errors.ErrEmptyBooks)
	}

	var bookPrint []string

	for _, book := range postRequestBook.Books {
		bookPrint = append(bookPrint, book)
	}

	postRequestBook.booksListString = strings.Join(bookPrint, ", ")

	logger.Logger.Debugf("get books called [books]: %s", postRequestBook.booksListString)

	return c.String(http.StatusOK, "Books: "+postRequestBook.booksListString)
}

func PostBooks(c echo.Context) error {

	if err := c.Bind(&postRequestBook); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.ErrUnmarshalFail)
	}

	if postRequestBook.Books == nil {
		return c.JSON(http.StatusBadRequest, errors.ErrEmptyBooks)
	}

	var bookPrint []string

	for _, book := range postRequestBook.Books {
		bookPrint = append(bookPrint, book)
	}

	postRequestBook.booksListString = strings.Join(bookPrint, ", ")

	logger.Logger.Debugf("post books called [books]: %s", postRequestBook.booksListString)

	logger.Logger.Info("books successfully added!")

	return c.String(http.StatusOK, fmt.Sprintf("[Books]: %s successfully added!", postRequestBook.booksListString))
}

func PutBooks(c echo.Context) error {

	if err := c.Bind(&putRequestBook); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.ErrUnmarshalFail)
	}

	if postRequestBook.Books == nil {
		return c.JSON(http.StatusBadRequest, errors.ErrEmptyBooks)
	}

	var bookPrint []string

	for _, book := range putRequestBook.NewBooks {
		bookPrint = append(bookPrint, book)
	}

	postRequestBook.Books = putRequestBook.NewBooks

	putRequestBook.newBooksListString = strings.Join(bookPrint, ", ")

	logger.Logger.Debugf("put books called [books]: %s", postRequestBook.booksListString)

	logger.Logger.Info("books successfully updated!")

	return c.String(http.StatusOK, fmt.Sprintf("Book successfully updated, [Old Books]: %s, [New Books]: %s", postRequestBook.booksListString, putRequestBook.newBooksListString))
}

func DeleteBooks(c echo.Context) error {

	if postRequestBook.Books == nil {
		return c.JSON(http.StatusBadRequest, errors.ErrEmptyBooks)
	}

	deleteBook := c.QueryParam("book")

	if deleteBook == "all" {
		postRequestBook.Books = nil
		return c.String(http.StatusOK, "All books was deleted successfully!")
	}

	index, err := strconv.Atoi(deleteBook)
	if err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "server error, cannot convert index",
		})
	}

	index--

	if index >= 0 && index < len(postRequestBook.Books) {
		postRequestBook.Books = append(postRequestBook.Books[:index], postRequestBook.Books[index+1:]...)
	}

	logger.Logger.Debugf("delete books called [book index]: %s", deleteBook)

	logger.Logger.Info("books successfully deleted!")

	return c.String(http.StatusOK, "Book was deleted!")
}
