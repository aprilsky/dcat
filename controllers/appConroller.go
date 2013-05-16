package controllers

import (
	"dcat/models"
	"dcat/service"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func saveApp(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	data := req.FormValue("data")
	if data != "" {
		b := []byte(data)
		appinfo := models.NewAppInfo()
		json.Unmarshal(b, appinfo)
		errMsg, err := service.Addapp(appinfo)
		fmt.Println(errMsg, "====", err)
		rw.Write(b)
	}
}

func getAppUserinfo() {

}

func NewappHandler(rw http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		fmt.Println("=======2=========")
		t, _ := template.ParseFiles("views/app.html")
		t.Execute(rw, nil)
	} else {
		saveApp(rw, req)
	}

}
