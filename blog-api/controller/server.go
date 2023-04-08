package controller

import (
	"fmt"
	"log"
	"net/http"

	controller "blog-api/controller/blog"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}