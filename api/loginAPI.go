package api

import (
	"dcat/service"
	"encoding/json"
	"log"
	"net/http"
)

func LoginHandlerAPI(rw http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		req.ParseForm()
		name := req.FormValue("name")
		pwd := req.FormValue("password")
		log.Printf("name", name, pwd)
		account := service.Login(name, pwd)
		log.Printf("account", account.AccountName, account.Pwd)
		if account.Pwd == pwd {
			a, _ := json.Marshal(account)
			b := []byte(a)
			rw.Write(b)
		}
	}
}
