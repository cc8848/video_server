package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"encoding/json"
	"io"
	"io/ioutil"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := &HomePage{
		Name: "Zereker",
	}
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		t, err := template.ParseFiles("./templates/home.html")
		if err != nil {
			log.Printf("Parsing templates home.html error: %v\n", err)
			return
		}
		t.Execute(w, p)
		return
	}
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("username")

	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{
			Name: cname.Value,
		}
	} else if len(fname) != 0 {
		p = &UserPage{
			Name: fname,
		}
	}

	t, err := template.ParseFiles("./templates/userhome.html")
	if err != nil {
		log.Printf("Parse userhome.html error: ", err)
	}
	t.Execute(w, p)
}

func apiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(RequestNotRecognizedError)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		re, _ := json.Marshal(RequestBodyParseFailedError)
		io.WriteString(w, string(re))
		return
	}

	request(apiBody, w, r)
	defer r.Body.Close()
}
