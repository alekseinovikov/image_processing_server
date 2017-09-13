package main

import (
	"fmt"
	"flag"
	"net/http"
	_ "github.com/freedom4live/resize"
)


// Stores configs of the server
type ServerConfigs struct {
	port string
}


// Parses command line arguments and puts them into the ServerConfigs structure
func parseArgsToConfigs() ServerConfigs {
	portPtr := flag.String("port", "8080", "The portPtr of the server")
	flag.Parse()

	return ServerConfigs{port: *portPtr}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

//Entry point of the program
func main() {
	configs:= parseArgsToConfigs()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":" + configs.port, nil)
}
