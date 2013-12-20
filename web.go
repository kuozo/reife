package main

import (
	"fmt"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln("Hi, The App comming soon")
}

func main() {
	http.HandleFunc("/", HomeHandler)
	fmt.Println("Listening.........")
	err := http.ListenAndServe(":"+os.Getenv("Port"), nil)
	if err != nil {
		panic(err)
	}
}
