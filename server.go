package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func getMyIpAddress(w http.ResponseWriter, r *http.Request)  {
	url := "http://ipinfo.io"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray)) // htmlをstringで取得
}

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	fmt.Print(" \n ---Hello World---\n")
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/ip", helloWorld)
	http.ListenAndServe(":8080", nil)
}
