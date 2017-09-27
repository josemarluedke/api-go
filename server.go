package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/josemarluedke/api/config"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	fmt.Printf("Server listening on port %s", config.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Config.Port), nil))
}
