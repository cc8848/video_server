package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("user_name")
	io.WriteString(w, username)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
