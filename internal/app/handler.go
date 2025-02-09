package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler represents the handler for the API
type Handler struct {
	service *Service
}

// InitHandler initializes the handler component
func InitHandler(service *Service) *Handler {
	return &Handler{service}
}

// Ping hadnles the /ping endpoint and responds with "Pong!" to indicate that the server is running
func (handler *Handler) Ping(writer http.ResponseWriter, request *http.Request) {
	// Validating the request method, should be GET
	handleHttpMethod(writer, request, http.MethodGet)
	handleSuccessResponse(writer, "Pong!", nil)
}

// Create handles the /create endpoint and processes the creation of a short link.
func (handler *Handler) Create(writter http.ResponseWriter, request *http.Request) {
	// Validating the request method, should be POST
	handleHttpMethod(writter, request, http.MethodPost)

	// CreateShortLinkRequest struct to hold the request payload
	var payload CreateShortLinkRequest
	handlePayload(writter, request, &payload)

	// CreateMasterUrl method from the service layer
	response, error := handler.service.CreateMasterUrl(payload)
	if error != nil {
		handleErrorResponse(writter, http.StatusBadRequest, error.Error())
		return
	}

	handleSuccessResponse(writter, "Short link created successfully", response)
}

// Redirect handles the /redirect endpoint and processes the redirection to full url based on the key from short url.
func (handler *Handler) Redirect(writter http.ResponseWriter, request *http.Request) {
	// Validating the request method, should be GET
	handleHttpMethod(writter, request, http.MethodGet)
	vars := mux.Vars(request)
	shortKey := vars["key"]

	// GetTargetUrl method from the service layer
	targetUrl, error := handler.service.GetTargetUrl(shortKey)
	// Handle and return error if any
	if error != nil {
		handleErrorResponse(writter, http.StatusNotFound, error.Error())
		return
	}

	// Redirect to target url
	http.Redirect(writter, request, targetUrl, http.StatusMovedPermanently)
}

// handleSuccessResponse sends a successful JSON response with the given message and data.
func handleSuccessResponse(writer http.ResponseWriter, message string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(ApiResponse{
		Status:  "Success",
		Message: message,
		Data:    data,
	})
}

// handleErrorResponse sends an error JSON response with the given status and message.
func handleErrorResponse(writer http.ResponseWriter, status int, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(ApiResponse{
		Status:  "error",
		Message: message,
	})
}

// handleHttpMethod checks if the request method matches the expected method.
func handleHttpMethod(writer http.ResponseWriter, request *http.Request, requestedMethod string) {
	if request.Method != requestedMethod {
		http.Error(writer, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
}

// handlePayload decodes the JSON payload from the request body into the provided payload struct.
func handlePayload(writer http.ResponseWriter, request *http.Request, payload interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&payload)
	fmt.Printf("Request Body: %+v\n", payload)
	if err != nil {
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}
}
