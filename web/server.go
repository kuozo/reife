package main

import (
	"fmt"
	"http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, The app comming soon.")
}

func main() {
	http.HandleFunc("/", HomeHandler)
	err := http.ListenAndService(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
