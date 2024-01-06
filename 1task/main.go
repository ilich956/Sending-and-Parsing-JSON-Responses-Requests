package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	Address      Address   `json:"address"`
	Contacts     []Contact `json:"contacts"`
	IsStudent    bool      `json:"isStudent"`
	Grades       []int     `json:"grades"`
	RegisteredAt string    `json:"registeredAt"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Contact struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type RequestData struct {
	Person Person `json:"person"`
	Status string `json:"status"`
}

type ResponseData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const port = ":8080"

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server listening on port", port)
	http.ListenAndServe(port, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlePostRequest(w, r)
	case http.MethodGet:
		handleGetRequest(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	var requestData RequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		handleErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validate the received data or perform any necessary actions
	// ...

	response := ResponseData{
		Status:  http.StatusOK,
		Message: "Data successfully received",
	}

	sendJSONResponse(w, response, http.StatusOK)
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	// In a real application, you might perform some actions based on the GET request
	// For this example, let's just send a sample response
	response := ResponseData{
		Status:  http.StatusOK,
		Message: "GET request received",
	}

	sendJSONResponse(w, response, http.StatusOK)
}

func handleErrorResponse(w http.ResponseWriter, status int, message string) {
	response := ResponseData{
		Status:  status,
		Message: message,
	}

	sendJSONResponse(w, response, status)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	responseJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
