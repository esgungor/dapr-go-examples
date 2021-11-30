package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Success string `json:"success"`
}

func getSecret(w http.ResponseWriter, r *http.Request) {
	sUrl := "http://localhost:3500/v1.0/secrets/secretstore/basic-auth"
	resp, err := http.Get(sUrl)
	if err != nil {
		fmt.Println(err)
	}
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	w.Write(data)

}

func sender(w http.ResponseWriter, r *http.Request) {
	publishUrl := `http://localhost:3500/v1.0/publish/pubsub/deathStarStatus`
	// msg := map[string]string{
	// 	"data": "test",
	// }
	msg := Message{
		Success: "hello",
	}
	enc, err := json.Marshal(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enc))
	req, err := http.Post(publishUrl, "application/json", bytes.NewBuffer(enc))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(req.Status)
	// var buffer = make([]byte, 32*1024)
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	w.Write(res)

	fmt.Println("Sended!", string(res))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", sender).Methods("POST")
	router.HandleFunc("/secret", getSecret).Methods("GET")

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
