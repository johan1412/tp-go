package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"bufio"
)

func currentTime(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
			now := time.Now()
			heure := now.Format("15:04")
			fmt.Fprintf(w, heure)
		case http.MethodPost:
			fmt.Fprintf(w, "Bad method")
	}
}

func save(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
			fmt.Println("Bad method")
		case http.MethodPost:
			if err := req.ParseForm(); err != nil {
				fmt.Println("Something went bad")
				fmt.Fprintln(w, "Something went bad")
				return
			}
			author := req.PostForm["author"][0]
			entry := req.PostForm["entry"][0]
			saveFile, err := os.OpenFile("./entries.data", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
			defer saveFile.Close()

			writter := bufio.NewWriter(saveFile)
			if err == nil {
				fmt.Fprintf(writter, "%s:%s\n", author, entry)
			}
			writter.Flush()

			fmt.Fprintf(w, "%s : %s\n", author, entry)
	}
}

func entries(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
			
		case http.MethodPost:
		
	}
}

func main() {
	http.HandleFunc("/", currentTime)
	http.HandleFunc("/add", save)
	http.HandleFunc("/entries", entries)
	http.ListenAndServe(":4567", nil)
}