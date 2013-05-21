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
	log.Printf(data)
	if data != "" {
		b := []byte(data)
		appinfo := new(models.AppInfo)
		json.Unmarshal(b, appinfo)
		errMsg, err := service.Addapp(appinfo)
		if err != nil {
			fmt.Println(errMsg, "====", err)
		}
		rw.Write(b)
	}
}

/** 
 * 查看详细APP信息
 */
func ViewAppInfoHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Println(req.FormValue("id"))
	t, _ := template.ParseFiles("views/appdetail.html")
	t.Execute(rw, nil)
}

/**
 * 打开新增APP 页面且保存APP
 */
func NewappHandler(rw http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		log.Println("=======2=========")
		t, _ := template.ParseFiles("views/app.html")
		t.Execute(rw, nil)
	} else {
		saveApp(rw, req)
	}

}

func GetAppListHandler(rw http.ResponseWriter, req *http.Request) {
	sess := GlobalSessions.SessionStart(rw, req)
	var account *models.AccountInfo
	if req.Method == "POST" {
		accountSess := sess.Get("account")
		if accountSess != nil {
			account = accountSess.(*models.AccountInfo)
			applist := service.GetAppUserinfo(account.Aid)
			a, _ := json.Marshal(applist)
			b := []byte(a)
			rw.Write(b)
		}
	}

}
