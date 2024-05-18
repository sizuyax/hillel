package http_server

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

const port = "8081"

func HttpServer() {
	http.Handle("/", http.HandlerFunc(handle))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		logrus.Fatal(err)
	}

	logrus.Printf("server is listening on port %s HTTP Server", port)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handleGet(w, r)
	}

	if r.Method == http.MethodPost {
		handlePost(w, r)
	}

	if r.Method == http.MethodPut {
		handlePut(w, r)
	}

	if r.Method == http.MethodDelete {
		handleDelete(w, r)
	}
}

type handlePostRequest struct {
	Movie string
	Grade int
}

type handlePutRequest struct {
	Comment    int
	NewComment string
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		w.Write([]byte("failed to read body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logrus.Println(string(reqBytes))

	query := r.URL.Query()

	nameValue := query.Get("name")
	ageValue := query.Get("age")

	age, err := strconv.Atoi(ageValue)
	if err != nil {
		logrus.Error(err)
		w.Write([]byte("failed to convert age"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.Println("GET params: ", "[name]: "+nameValue, "[age]: ", age)

	w.Write([]byte(fmt.Sprintf("query param [name]: %s, [age]: %v", nameValue, age)))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	postRequest := &handlePostRequest{}

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		w.Write([]byte("failed to read body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(reqBytes, postRequest); err != nil {
		logrus.Error(err)
		w.Write([]byte("failed unmarshal request"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if postRequest.Grade > 5 {
		logrus.Error(err)
		w.Write([]byte("grade can't be bigger than 5"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.Printf("[Movie]: %s   [Grade]: %v", postRequest.Movie, postRequest.Grade)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	putRequest := &handlePutRequest{}

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		w.Write([]byte("failed to read body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(reqBytes, putRequest); err != nil {
		logrus.Error(err)
		w.Write([]byte("failed unmarshal request"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.Printf("[Comment]: %v   [New Comment]: %s", putRequest.Comment, putRequest.NewComment)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("firstname")
	lastname := r.URL.Query().Get("lastname")

	logrus.Println("DELETE params: ", "[firstname]: "+firstname, " [lastname]: "+lastname)

	w.Write([]byte(fmt.Sprintf("query param [firstname]: %s,  [lastname]: %s", firstname, lastname)))
}
