package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Reading web pages in Go.")
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error in reading the webpage")
		log.Fatal("Error : ", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
