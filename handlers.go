package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var imageFolder string = "img/"
var imageExtension string = ".png"

func getPath(name string) string {
	return imageFolder + name + imageExtension
}

func GetImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filename := ps.ByName("uuid")
	fileContent, err := os.ReadFile(getPath(filename))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileContent)
}

func PostImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseMultipartForm(32 << 10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	fileContent, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
		return
	}

	filename := ps.ByName("uuid")
	file, err := os.Create(getPath(filename))
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
		return
	}

	_, err = io.Copy(file, fileContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
		return
	}

	w.WriteHeader(200)

	log.Printf("Created file %s", filename)
}

func DeleteImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filename := ps.ByName("uuid")
	filePath := getPath(filename)
	_, err := os.Stat(filePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
		return
	}

	w.WriteHeader(200)

	log.Printf("Deleted file %s", filename)
}
