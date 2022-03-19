package main

import (
	"log"
	"net/http"
  "encoding/json"
)

func main() {

	http.HandleFunc("/myendpoint", AllowCORS(myHandler))

	log.Fatal(http.ListenAndServe(":9000", nil))
}

  
func myHandler(w http.ResponseWriter, r *http.Request) {
  type response struct {
    Message string
  }

  json.NewEncoder(w).Encode(response{Message: "Here is a message"})
}

func AllowCORS(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		handler.ServeHTTP(w, r)
	})
}


