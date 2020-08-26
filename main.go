package main

import (
	"github.com/yigitsadic/qrmenum_client/client"
	"github.com/yigitsadic/qrmenum_client/handlers"
	"github.com/yigitsadic/qrmenum_client/store"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cmsUrl := os.Getenv("CMS_URL")
	if cmsUrl == "" {
		cmsUrl = "http://localhost:5000"
	}

	log.Println("Starting server on PORT", port)

	c := client.NewHTTPClient(cmsUrl)
	s := store.NewInMemoryStore(c)

	http.HandleFunc("/", handlers.ProductHandler(s))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln("Unable to start server on port", port)
	}
}
