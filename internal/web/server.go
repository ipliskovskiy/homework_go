package web

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
)

type StartInfo struct{
	Title string
	Message string
	BottomInformation string
}

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		data := StartInfo {
			Title: "TicTacToe",
			Message: "Добро пожаловать в увлекательную интерактивную игру крестики нолики.",
			BottomInformation: "Version: Experemental",
		}
		tmpl, _ := template.ParseFiles("internal/web/templates/index.html")
        tmpl.Execute(w, data)
	})

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("internal/web/templates/css"))))
	http.HandleFunc("/hello", sayhello)
	fmt.Println("Start server...")
	http.ListenAndServe(":8080", nil)
}


func sayhello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Привет!")
}