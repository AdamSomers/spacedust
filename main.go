package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("hello, asdf\n")
	doSomething()
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", fs)

	//err := http.ListenAndServe(":8000", nil)
	err := http.ListenAndServeTLS(":443", "cert/adamsomers_com.crt", "cert/adamsomers.com.key", nil)
	log.Fatal(err)
	fmt.Println("bye")
}

func doSomething() {
	fmt.Printf("I am doing it!\n")
}
