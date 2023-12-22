package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var request JsonRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if request.Message == "" {
		errorResponse := JsonResponse{
			Status:  "400",
			Message: "Invalid JSON message",
		}

		errorResponseJSON, err := json.Marshal(errorResponse)
		if err != nil {
			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorResponseJSON)
		return
	}

	fmt.Printf("Received message: %s\n", request.Message)

	response := JsonResponse{
		Status:  "success",
		Message: "Data successfully received",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
