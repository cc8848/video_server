package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/url"
	"net/http/httputil"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)

	router.POST("/upload/:vid-id", proxyHandler)
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))
	return router
}

func proxyHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u, _ := url.Parse("http://127.0.0.1:9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.ListenAndServe(":8080", RegisterHandler())
}
