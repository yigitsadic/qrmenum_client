package main

import (
	"github.com/yigitsadic/qrmenum_client/client"
	"github.com/yigitsadic/qrmenum_client/handlers"
	"github.com/yigitsadic/qrmenum_client/store"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cmsUrl := os.Getenv("CMS_URL")
	if cmsUrl == "" {
		cmsUrl = "http://localhost:5000"
	}

	showTmpl := template.Must(template.ParseFiles("templates/show.html"))
	notFoundTmpl := template.Must(template.ParseFiles("templates/404.html"))

	log.Println("Starting server on PORT", port)

	c := client.NewHTTPClient(cmsUrl)
	s := store.NewInMemoryStore(c)

	http.HandleFunc("/", handlers.ProductHandler(s, &mu, showTmpl, notFoundTmpl))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln("Unable to start server on port", port)
	}
}
