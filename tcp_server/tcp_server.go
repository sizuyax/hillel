package tcp_server

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func TcpServer() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		logrus.Fatal("error:", err)
		return
	}

	defer listener.Close()

	logrus.Println("Server is listening on port 8081 TCP Server")

	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Error("error:", err)
			continue
		}

		go handleClient(conn)
	}
}

type handlePostRequest struct {
	Movie string `json:"movie"`
	Grade int    `json:"grade"`
}

type handlePutRequest struct {
	Comment    int    `json:"comment"`
	NewComment string `json:"newComment"`
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				continue
			}
			logrus.Error("error reading:", err)
			return
		}

		request := string(buffer[:n])

		parts := strings.Split(request, " ")

		if len(parts) < 2 || parts[0] == "GET" {

			responseToUser := map[string]string{}

			statusCode := 0

			path := parts[1]

			pathParts := strings.Split(path, "?")
			if len(pathParts) < 2 {
				continue
			}

			queryString := pathParts[1]

			queryParams := strings.Split(queryString, "&")

			var nameKey string
			var nameValue string
			var ageKey string
			var ageValue int

			for _, param := range queryParams {
				paramParts := strings.Split(param, "=")

				key := paramParts[0]
				value := paramParts[1]

				if key == "name" {
					nameKey = key
					nameValue = value
				}

				if key == "age" {
					ageKey = key

					ageValue, err = strconv.Atoi(value)
					if err != nil {
						logrus.Error(err)
						responseToUser["error"] = "failed to convert age"
						statusCode = http.StatusInternalServerError
						break
					}
				}
			}

			if statusCode == 0 {
				responseToUser["message"] = "OK"
				statusCode = http.StatusOK
			}

			statusText := "OK"
			if statusCode != http.StatusOK {
				statusText = http.StatusText(statusCode)
			}

			jsonToUser, err := json.Marshal(responseToUser)
			if err != nil {
				logrus.Error(err)
				conn.Write([]byte("failed to marshal response"))
			}

			httpResponse := fmt.Sprintf("HTTP/1.1 %v %s\r\nContent-Length: %d\r\nContent-Type: application/json\r\n\r\n%s", statusCode, statusText, len(jsonToUser), string(jsonToUser))

			conn.Write([]byte(httpResponse))

			if statusCode != http.StatusOK {
				return
			}

			logrus.Printf("[%s]: %s, [%s]: %v", nameKey, nameValue, ageKey, ageValue)

		}

		if len(parts) < 2 || parts[0] == "POST" {

			responseToUser := map[string]string{}

			statusCode := 0

			bodyStart := strings.Index(request, "\r\n\r\n")
			if bodyStart == -1 {
				logrus.Error("invalid HTTP request")
				responseToUser["error"] = "invalid HTTP request"
				statusCode = http.StatusBadRequest
			}

			bodyStart += 4

			body := request[bodyStart:]

			var postReq handlePostRequest

			if err := json.Unmarshal([]byte(body), &postReq); err != nil {
				logrus.Error("error decoding JSON:", err)
				responseToUser["error"] = "error decoding JSON"
				statusCode = http.StatusInternalServerError
			}

			if postReq.Grade > 5 {
				responseToUser["error"] = "grade can't be bigger than 5"
				statusCode = http.StatusBadRequest
			}

			if statusCode == 0 {
				responseToUser["message"] = "OK"
				statusCode = http.StatusOK
			}

			statusText := "OK"
			if statusCode != http.StatusOK {
				statusText = http.StatusText(statusCode)
			}

			jsonToUser, err := json.Marshal(responseToUser)
			if err != nil {
				logrus.Error(err)
				conn.Write([]byte("failed to marshal response"))
			}

			httpResponse := fmt.Sprintf("HTTP/1.1 %v %s\r\nContent-Length: %d\r\nContent-Type: application/json\r\n\r\n%s", statusCode, statusText, len(jsonToUser), string(jsonToUser))

			conn.Write([]byte(httpResponse))

			if statusCode != http.StatusOK {
				return
			}

			logrus.Printf("[movie]: %s, [grade]: %v", postReq.Movie, postReq.Grade)
		}

		if len(parts) < 2 || parts[0] == "PUT" {

			responseToUser := map[string]string{}

			statusCode := 0

			bodyStart := strings.Index(request, "\r\n\r\n")
			if bodyStart == -1 {
				logrus.Error("invalid HTTP request")
				responseToUser["error"] = "invalid HTTP request"
				statusCode = http.StatusBadRequest
			}

			bodyStart += 4

			body := request[bodyStart:]

			var putReq handlePutRequest

			if err := json.Unmarshal([]byte(body), &putReq); err != nil {
				logrus.Error("error decoding JSON:", err)
				responseToUser["error"] = "error decoding JSON"
				statusCode = http.StatusInternalServerError
			}

			if statusCode == 0 {
				responseToUser["message"] = "OK"
				statusCode = http.StatusOK
			}

			statusText := "OK"
			if statusCode != http.StatusOK {
				statusText = http.StatusText(statusCode)
			}

			jsonToUser, err := json.Marshal(responseToUser)
			if err != nil {
				logrus.Error(err)
				conn.Write([]byte("failed to marshal response"))
			}

			httpResponse := fmt.Sprintf("HTTP/1.1 %v %s\r\nContent-Length: %d\r\nContent-Type: application/json\r\n\r\n%s", statusCode, statusText, len(jsonToUser), string(jsonToUser))

			conn.Write([]byte(httpResponse))

			if statusCode != http.StatusOK {
				return
			}

			logrus.Printf("[comment]: %v, [new comment]: %v", putReq.Comment, putReq.NewComment)
		}

		if len(parts) < 2 || parts[0] == "DELETE" {

			responseToUser := map[string]string{}

			statusCode := 0

			path := parts[1]

			pathParts := strings.Split(path, "?")
			if len(pathParts) < 2 {
				continue
			}

			queryString := pathParts[1]

			queryParams := strings.Split(queryString, "&")

			var firstnameKey string
			var firstnameValue string
			var lastnameKey string
			var lastnameValue string

			for _, param := range queryParams {
				paramParts := strings.Split(param, "=")

				key := paramParts[0]
				value := paramParts[1]

				if key == "firstname" {
					firstnameKey = key
					firstnameValue = value
				}

				if key == "lastname" {
					lastnameKey = key
					lastnameValue = value
				}
			}

			if statusCode == 0 {
				responseToUser["message"] = "OK"
				statusCode = http.StatusOK
			}

			statusText := "OK"
			if statusCode != http.StatusOK {
				statusText = http.StatusText(statusCode)
			}

			jsonToUser, err := json.Marshal(responseToUser)
			if err != nil {
				logrus.Error(err)
				conn.Write([]byte("failed to marshal response"))
			}

			httpResponse := fmt.Sprintf("HTTP/1.1 %v %s\r\nContent-Length: %d\r\nContent-Type: application/json\r\n\r\n%s", statusCode, statusText, len(jsonToUser), string(jsonToUser))

			conn.Write([]byte(httpResponse))

			if statusCode != http.StatusOK {
				return
			}

			logrus.Printf("[%s]: %s, [%s]: %v", firstnameKey, firstnameValue, lastnameKey, lastnameValue)

		}
	}

}
