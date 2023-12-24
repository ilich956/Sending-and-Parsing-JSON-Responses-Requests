package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const port = ":8080"

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server listening on port ", port)
	http.ListenAndServe(port, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	var requestData RequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if requestData.Message == "" {
		response := ResponseData{
			Status:  "400",
			Message: "Invalid JSON message",
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJSON)
		return
	}

	fmt.Printf("Received message: %s\n", requestData.Message)

	response := ResponseData{
		Status:  "success",
		Message: "Data successfully received",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
