package web

import (
	"Ex2_Week3/pkg/chatgpt" // Make sure the import path matches your project structure
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, client *chatgpt.Client) {
	router.HandleFunc("/ask", AskHandler(client)).Methods("POST")                                            // Make sure this is correct
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // Serving static files
	router.HandleFunc("/", serveHome).Methods("GET")                                                         // Serve your home page on the root path
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
