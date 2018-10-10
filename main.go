package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Go is Awesome!</h1>")
}

func main(){
	http.HandleFunc("/",indexHandler)
	http.ListenAndServe(":8082",nil)
}
