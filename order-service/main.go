package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func listener(w http.ResponseWriter, r *http.Request) {
	resp, _ := io.ReadAll(r.Body)
	fmt.Println("Triggered:", string(resp))
	w.Write(resp)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/listen", listener).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))
}
