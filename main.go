package main

import (
	"fmt"
	"log"
	"net/http"
)


func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	fmt.Printf("Hello Go");
}

func main() {
	http.HandleFunc("/", sayHello)
	http.Handle("/home/", http.StripPrefix("/home", http.FileServer(http.Dir("public"))))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}