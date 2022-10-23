package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"auteur":"Djibril Thiongane","profession":"future talent01","app": "a minimal server in golang"}`))
}

func main() {
	fmt.Println("vim-go")
	fmt.Println("what is appening")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}
