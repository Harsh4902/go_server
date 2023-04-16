package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHendler)
	http.HandleFunc("/hello",helloHendler)

	log.Fatal(http.ListenAndServe(":8080",nil))
}

func helloHendler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Hello from my Servre</h1>"))
}

func formHendler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Fprintf(w,"ParseForm() error: %v",err)
		return
	}

	fmt.Fprintf(w,"Post request successful\n")
	fmt.Fprintf(w,"Name: %v\n",r.FormValue("name"))
	fmt.Fprintf(w,"Name: %v\n",r.FormValue("address"))
}