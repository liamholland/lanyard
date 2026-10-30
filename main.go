package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "watta nuba pon para tuba pre dunkieeee")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Srcrib dooooonnnnnn")
}
