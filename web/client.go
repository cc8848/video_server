package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"io"
	"bytes"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(apiBody *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error

	switch apiBody.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", apiBody.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("error : %v", err)
			return
		}
		normalResponse(w, resp)
	case http.MethodPost:
		req, _ := http.NewRequest("POST", apiBody.Url, bytes.NewBuffer([]byte(apiBody.ReqBody)))
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("error : %v", err)
			return
		}
		normalResponse(w, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("DELETE", apiBody.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("error : %v", err)
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request")
		return
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(InternalFaultsError)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(string(re)))
		return
	}
	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}
