package main

import (
	"net/http"
	"os"
	"time"
	"io/ioutil"
	"log"
	"io"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	videoPath := VideoDir + vid
	video, err := os.Open(videoPath)
	if err != nil {
		log.Printf("Error when try to open file: %v\n", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer video.Close()
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}
	file, _, err := r.FormFile("file") // <form name="file">
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v\n", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	vid := p.ByName("vid-id")
	videoPath := VideoDir + vid
	err = ioutil.WriteFile(videoPath, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}

func testPageHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}
