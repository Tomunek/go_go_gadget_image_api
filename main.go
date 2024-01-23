package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var banner string = ` ___                               _    ____ ___ ` + "\n" +
	`|_ _|_ __ ___   __ _  __ _  ___   / \  |  _ \_ _|` + "\n" +
	` | ||  _   _ \ / _' |/ _  |/ _ \ / _ \ | |_) | | ` + "\n" +
	` | || | | | | | (_| | (_| |  __// ___ \|  __/| | ` + "\n" +
	`|___|_| |_| |_|\__,_|\__, |\___/_/   \_\_|  |___|` + "\n" +
	`                     |___/                       `

func HealthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{status: ok}"))
}

func main() {
	fmt.Println(banner)
	router := httprouter.New()
	router.GET("/", HealthHandler)
	router.GET("/img/*uuid", GetImageHandler)
	router.POST("/img/*uuid", PostImageHandler)
	router.DELETE("/img/*uuid", DeleteImageHandler)
	log.Println("Starting image service")

	log.Fatal(http.ListenAndServe(":8080", router))
}
