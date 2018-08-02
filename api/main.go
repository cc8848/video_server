package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	newMiddleWareHandler := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", newMiddleWareHandler)
}
