package main

import (
	"fmt"
	"net/http"
	"os"
)

//localhost:4000/  returns homepage
//localhost:4000/shutdown turn off app
func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":4000", nil)
}
func shutdown(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Home Page")
}
