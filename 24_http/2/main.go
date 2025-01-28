package main

import "net/http"

func main() {

	W:= http.NewServeMux()
	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Jay shree Ram"))
	})


	

}


