package controllers

import (
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello This is your request")
}
