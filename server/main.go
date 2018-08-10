package main

import (
	"fmt"
    "net/http"
)

func user(w http.ResponseWriter, req *http.Request) {
	userId := req.URL.Query().Get("userId")
	fmt.Printf("==== %v", str)
	w.Write([]byte("Hello"))
}
func say(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
    http.HandleFunc("/user/", user)
    http.HandleFunc("/login/", say)

    http.ListenAndServe(":8001", nil)
}
