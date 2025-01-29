package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Thsi is '/' route"))
	})

	http.HandleFunc("GET /api/get/users/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("inside the GET /api/get/users/")
		fmt.Printf("Method is : %v", req.Method)
		if req.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Wring CRUD HTTp Method"))
			return
		}
		w.Write([]byte("This is 'GET /api/get/users/'"))
	})

	http.ListenAndServe(":8055", nil)
}
