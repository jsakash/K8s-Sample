package main

import (
	"html/template"
	"log"
	"net/http"
)

// func init() {
// 	gin.SetMode(gin.ReleaseMode)
// }
var tpl *template.Template

func main() {

	tpl, _ = tpl.ParseGlob("home.html")
	http.HandleFunc("/", home)
	log.Println("Web server has started..")
	http.ListenAndServe(":80", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)

}
