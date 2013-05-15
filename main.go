package main

import (
	"dcat/api"
	"dcat/controllers"
	"net/http"
)

func main() {

	http.Handle("/css/", http.FileServer(http.Dir("views")))
	http.Handle("/js/", http.FileServer(http.Dir("views")))
	http.Handle("/img/", http.FileServer(http.Dir("views")))

	http.HandleFunc("/add", controllers.NewappHandler)
	http.HandleFunc("/app/add", controllers.NewappHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/", controllers.IndexHandler)

	http.HandleFunc("/v1/login", api.LoginHandlerAPI)

	http.ListenAndServe("127.0.0.1:8080", nil)

}
