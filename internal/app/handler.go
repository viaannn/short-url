package app

import (
	"fmt"
	"net/http"
)

func HomeHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "Hello!")
}
