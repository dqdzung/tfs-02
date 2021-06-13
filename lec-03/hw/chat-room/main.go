package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var messages []map[string]string

func handlePost(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]string)
	json.Unmarshal(reqBody, &m)

	messages = append(messages, m)
	data, _ := json.Marshal(messages)

	_ = ioutil.WriteFile("output.json", data, 0644)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("output.json")
	fmt.Fprintln(w, string(data))
}

func main() {
	log.Println("Server started on port 8080")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/chat", handlePost)
	http.HandleFunc("/message", handleGet)
	http.ListenAndServe(":8080", nil)
}
