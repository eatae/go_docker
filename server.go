package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myOS, myArch := runtime.GOOS, runtime.GOARCH
		inContainer := "inside"
		if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
			inContainer = "outside"
		}
		_, _ = fmt.Fprintf(w, "Hello, %s!\n", r.UserAgent())
		_, _ = fmt.Fprintf(w, "I'm running on %s/%s.\n", myOS, myArch)
		_, _ = fmt.Fprintf(w, "I'm - running %s of a container.\n", inContainer)
	})

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatalln(err)
	}
}
