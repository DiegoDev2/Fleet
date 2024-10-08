package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/arch/", root)
	http.ListenAndServe(":8080", mux)
	fmt.Printf("Listening on port 8080\n")

}
func root(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./public")).ServeHTTP(w, r)
}
