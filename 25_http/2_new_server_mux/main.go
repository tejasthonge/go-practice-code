package main

import (
	"fmt"
	"net/http"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "this is users list")
}

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Jay shree Ram"))
	})
	//creating separate handeler
	router.HandleFunc("GET /api/users/", getUserHandler)
	server := http.Server{
		Addr:    ":8055",
		Handler: router,
	}
	fmt.Println("Starting server")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting the server %v..", err)
	}
}
