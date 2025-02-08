package app

import (
	"fmt"
	"net/http"
)

type Handler struct {
	service *Service
}

func InitHandler(service *Service) *Handler {
	return &Handler{service}
}

func (handler *Handler) Ping(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(writer, "Pong!")
}

func (handler *Handler) Create(writter http.ResponseWriter, request *http.Request) {

}

func (handler *Handler) Redirect(writter http.ResponseWriter, request *http.Request) {

}
