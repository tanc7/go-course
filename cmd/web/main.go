package main

import (
	"github.com/tanc7/go-course/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

//Home is / path

//main wtf bro.
func main() {
	//fmt.Println("Hello world!")
	// This is basically a router.rb or route.py module
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	//http.HandleFunc("/divide", Divide)
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprintf(w, "Hello world!")
	//	if err != nil {
	//		//log.Fatal(err)
	//		fmt.Println(err)
	//	}
	//	//fmt.Sprintf("Bytes written:%d", n)
	//	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	//})
	_ = http.ListenAndServe(portNumber, nil)

}
