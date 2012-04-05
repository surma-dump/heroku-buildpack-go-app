package main

import (
	// "code.google.com/p/gorilla/mux"
	// testpkg "launchpad.net/heroku-buildpack-go-testpkg"
	"net/http"
	"os"
	"heroku-buildpack-go-app/consts"
	"fmt"
)

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/", handler)
	// e := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	// if e != nil {
		// panic(e)
	// }
	http.HandleFunc("/", handler)
	e := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if e != nil {
		panic(e)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, consts.MESSAGE+"\n"+testpkg.GiveMeMyString())
	fmt.Fprintf(w, consts.MESSAGE)
}
