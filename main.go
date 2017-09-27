package main

import (
	"fmt"
	"flag"
	"net/http"
	_ "github.com/freedom4live/resize"
	"strings"
)

// Stores configs of the server
type ServerConfigs struct {
	port string
}

const MESSAGE_METHOD_NOT_ALLOWED = "Method is not allowed! Supports POST only!"

// Parses command line arguments and puts them into the ServerConfigs structure
func parseArgsToConfigs() ServerConfigs {
	portPtr := flag.String("port", "8080", "The portPtr of the server")
	flag.Parse()

	return ServerConfigs{port: *portPtr}
}

// Handler for resizing image
func resizeHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(MESSAGE_METHOD_NOT_ALLOWED))

		return
	}

}

// Handler for creating thumbnail
func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(MESSAGE_METHOD_NOT_ALLOWED))

		return
	}

	fmt.Fprint(w, "Hello thumbnail")
}

//Entry point of the program
func main() {
	configs := parseArgsToConfigs()

	http.HandleFunc("/resize", resizeHandler)
	http.HandleFunc("/thumbnail", thumbnailHandler)
	http.ListenAndServe(":"+configs.port, nil)
}
