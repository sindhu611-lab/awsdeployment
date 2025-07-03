package main

import "net/http"

func msg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world:"))
}
func main() {
	router := http.NewServeMux()
	router.HandleFunc("/greet", msg)
	http.ListenAndServe(":8080", nil)
}
