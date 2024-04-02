package services

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
	Books      []string `json:"Books"`
	SplitBooks string
}

var putRequestBook struct {
	NewBooks      []string `json:"NewBooks"`
	SplitNewBooks string
}

var delRequestBook struct {
	DeleteBooks []string `json:"DeleteBooks"`
}

func GetBooks(c echo.Context) error {

	if postRequestBook.Books == nil {
		return c.JSON(http.StatusBadRequest, errors.Error{
			Code:    "INCORRECT_REQUEST",
			Message: "books cannot be empty!",
		})
	}

	var bookPrint []string

	for _, book := range postRequestBook.Books {
		bookPrint = append(bookPrint, book)
	}

	postRequestBook.SplitBooks = strings.Join(bookPrint, ", ")

	logger.Logger.Debugf("get books called [books]: %s", postRequestBook.SplitBooks)

	return c.String(http.StatusOK, "Books: "+postRequestBook.SplitBooks)
}

func PostBooks(c echo.Context) error {

	if err := c.Bind(&postRequestBook); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "INCORRECT_REQUEST",
			Message: "failed unmarshal request",
		})
	}

	if postRequestBook.Books == nil {
		return c.JSON(http.StatusBadRequest, errors.Error{
			Code:    "INCORRECT_REQUEST",
			Message: "books cannot be empty!",
		})
	}

	var bookPrint []string

	for _, book := range postRequestBook.Books {
		bookPrint = append(bookPrint, book)
	}

	postRequestBook.SplitBooks = strings.Join(bookPrint, ", ")

	logger.Logger.Debugf("post books called [books]: %s", postRequestBook.SplitBooks)

	logger.Logger.Info("books successfully added!")

	return c.String(http.StatusOK, fmt.Sprintf("[Books]: %s successfully added!", postRequestBook.SplitBooks))
}

func PutBooks(c echo.Context) error {

	if err := c.Bind(&putRequestBook); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "INCORRECT_REQUEST",
			Message: "failed unmarshal request",
		})
	}

	if postRequestBook.Books == nil {
		return c.JSON(http.StatusBadRequest, errors.Error{
			Code:    "INCORRECT_REQUEST",
			Message: "books cannot be empty!",
		})
	}

	var bookPrint []string

	for _, book := range putRequestBook.NewBooks {
		bookPrint = append(bookPrint, book)
	}

	postRequestBook.Books = putRequestBook.NewBooks

	putRequestBook.SplitNewBooks = strings.Join(bookPrint, ", ")

	logger.Logger.Debugf("put books called [books]: %s", postRequestBook.SplitBooks)

	logger.Logger.Info("books successfully updated!")

	return c.String(http.StatusOK, fmt.Sprintf("Book successfully updated, [Old Books]: %s, [New Books]: %s", postRequestBook.SplitBooks, putRequestBook.SplitNewBooks))
}

func DeleteBooks(c echo.Context) error {

	if err := c.Bind(&delRequestBook); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "INCORRECT_REQUEST",
			Message: "failed unmarshal request",
		})
	}

	if postRequestBook.Books == nil {
		return c.JSON(http.StatusBadRequest, errors.Error{
			Code:    "INCORRECT_REQUEST",
			Message: "books cannot be empty!",
		})
	}

	for _, book := range delRequestBook.DeleteBooks {
		if book == "all" {
			postRequestBook.Books = nil
			return c.String(http.StatusOK, "All books was deleted successfully!")
		}

		index, err := strconv.Atoi(book)
		if err != nil {
			logger.Logger.Error(err)
			return c.JSON(http.StatusInternalServerError, errors.Error{
				Code:    "INCORRECT_REQUEST",
				Message: "server error, cannot convert index",
			})
		}

		index--

		if index >= 0 && index < len(postRequestBook.Books) {
			postRequestBook.Books = append(postRequestBook.Books[:index], postRequestBook.Books[index+1:]...)
		}
	}

	logger.Logger.Debugf("delete books called [books]: %s", delRequestBook.DeleteBooks)

	logger.Logger.Info("books successfully deleted!")

	return c.String(http.StatusOK, "Book was deleted!")
}
