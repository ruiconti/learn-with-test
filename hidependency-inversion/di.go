package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	GreeterFD()
	log.Fatal(http.ListenAndServe(":1337", http.HandlerFunc(GreeterHandler)))
}

// A sockety approach
// Hint: http.ResponseWriter also implements io.Writer
func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Thompson :-)")
}

// A file descriptor approach
func GreeterFD() {
	Greet(os.Stdout, "Kernie")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
