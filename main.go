package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func currentTime(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
			now := time.Now()
			heure := now.Format("15:04")
			fmt.Println(heure)
		case http.MethodPost:
			fmt.Println("Bad method")
	}
}

func save(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
			fmt.Println("Bad method")
		case http.MethodPost:
			
		}
}

func main() {
	http.HandleFunc("/", currentTime)
	http.HandleFunc("/add", save)
	http.HandleFunc("/entries", helloHandler)
	http.ListenAndServe(":4567", nil)
}