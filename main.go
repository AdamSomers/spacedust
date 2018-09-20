package main

import (
	"fmt"
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

	http.ListenAndServe(":8000", nil)
	fmt.Println("bye")
}

func doSomething() {
	fmt.Printf("I am doing it!\n")
}
