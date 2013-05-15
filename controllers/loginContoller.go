package controllers

import (
	"dcat/models"
	"dcat/service"
	"html/template"
	"log"
	"net/http"
)

var account *models.AccountInfo

func LoginHandler(rw http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		t, err := template.ParseFiles("views/login.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(rw, nil)
	} else {
		req.ParseForm()
		name := req.FormValue("login")
		pwd := req.FormValue("password")
		log.Printf("name", name, pwd)
		account = service.Login(name, pwd)
		log.Printf("account", account.AccountName, account.Pwd)
		if account.Pwd == pwd {
			http.Redirect(rw, req, "/index", http.StatusFound)
		} else {

		}

	}

}
