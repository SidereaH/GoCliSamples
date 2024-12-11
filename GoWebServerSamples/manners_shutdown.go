package main

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
	"os"
	"os/signal"
)

// using manners for interrupt and kill server
func main() {
	handler := newHandler()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)
	manners.ListenAndServe(":4000", handler)

}
func newHandler() *handler {
	return &handler{}
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Andrey Hutornoy"
	}
	fmt.Fprintf(w, "Hello, my name is %s!", name)
}
func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}
