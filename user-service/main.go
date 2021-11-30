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

// var data = []byte("hello")
// var store = "my-store" // defined in the component YAML
type Store struct {
	Key   string `json:"key"`
	Value string `value:"key"`
}
type Stores []Store

func newState(w http.ResponseWriter, r *http.Request) {
	var test = Store{Key: "", Value: ""}
	var stores Stores
	json.NewDecoder(r.Body).Decode(&test)
	fmt.Println(test)
	stores = append(stores, test)
	sender, _ := json.Marshal(stores)
	fmt.Println(stores)
	fmt.Println(sender)

	result, err := http.Post("http://localhost:3500/v1.0/state/my-store", "application/json", bytes.NewBuffer(sender))
	if err != nil {
		panic(err)
	}

	testJ, _ := io.ReadAll(result.Body)
	fmt.Println(testJ)
	w.Write(testJ)
}

func getState(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	key := "engin"
	if q.Get("key") != "" {
		key = q.Get("key")
	}
	result, err := http.Get("http://localhost:3500/v1.0/state/my-store/" + key)
	if err != nil {
		panic(err)
	}
	testJ, _ := io.ReadAll(result.Body)
	fmt.Println(testJ)
	w.Write(testJ)
}
func main() {
	// ctx := context.Background()

	// state, err := client.GetState(ctx, store, "key1")
	router := mux.NewRouter()
	fmt.Println("HELLO MATE")
	router.HandleFunc("/add", newState).Methods("POST")
	router.HandleFunc("/", getState).Methods("GET")

	log.Fatal((http.ListenAndServe(":5000", router)))
}
