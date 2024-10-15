package main

import (
	"net/http"
	"fmt"
)
//1.Create web server

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to homepage")
	fmt.Println("Endpoint hit:homepage")
}
func main(){
	http.HandleFunc("/",homepage)
	http.ListenAndServe("localhost:10000",nil)
}