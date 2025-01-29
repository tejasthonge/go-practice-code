package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Jay Shree ram,Deafout")
	})

	http.HandleFunc("GET /api/users", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Thsi users api"))
	})
	http.ListenAndServe(":8055", nil)
}

//we can go any route like /api/sdfsf/ then also we geting the
/*
Jay Shree ram
*/
