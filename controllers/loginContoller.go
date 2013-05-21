package controllers

import (
	"dcat/service"
	"dcat/session"
	"html/template"
	"log"
	"net/http"
)

var GlobalSessions *session.SessionManager

//然后在init函数中初始化
func init() {
	GlobalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	log.Println(&GlobalSessions)
	go GlobalSessions.GC()
}

func LoginHandler(rw http.ResponseWriter, req *http.Request) {

	sess := GlobalSessions.SessionStart(rw, req)
	log.Println(&sess, sess.Get("account"))
	req.ParseForm()

	if req.Method == "GET" {
		t, err := template.ParseFiles("views/login.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(rw, sess.Get("account"))
	} else {

		name := req.FormValue("login")
		pwd := req.FormValue("password")
		log.Printf("name", name, pwd)
		account := service.Login(name, pwd)
		log.Printf("account", account.AccountName, account.Pwd)
		sess.Set("account", account)
		if account.Pwd == pwd {
			http.Redirect(rw, req, "/index", http.StatusFound)
		} else {

		}

	}

}

func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	sess := GlobalSessions.SessionStart(rw, req)
	log.Println(&sess, sess.Get("account"))
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Println(err)
	}

	account := sess.Get("account")
	t.Execute(rw, account)
}
