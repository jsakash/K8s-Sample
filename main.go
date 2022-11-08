package main

import (
	"log"
	"net/http"

	"github.com/jsakash/K8s-Sample/pkg/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Greet)
	http.HandleFunc("/home", controllers.Home)
	log.Println("Web server has started..")
	http.ListenAndServe(":8080", nil)
}
