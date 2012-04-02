package main

import (
	"code.google.com/p/gorilla/mux"
	"net/http"
	"os"
	"consts"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	e := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if e != nil {
		panic(e)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, consts.MESSAGE)
}
