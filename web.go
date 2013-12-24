package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", HomeHandler)
	fmt.Println("Listening.........")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
