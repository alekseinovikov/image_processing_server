package main

import (
	"flag"
	"net/http"
	"strings"
	"errors"
	"github.com/freedom4live/image_processing_server/resizing"
	"mime/multipart"
	"strconv"
)

// Stores configs of the server
type ServerConfigs struct {
	port string
}

const MESSAGE_METHOD_NOT_ALLOWED = "Method is not allowed! Supports POST only!"
const _24K = (1 << 10) * 24

// Parses command line arguments and puts them into the ServerConfigs structure
func parseArgsToConfigs() ServerConfigs {
	portPtr := flag.String("port", "8082", "The portPtr of the server")
	flag.Parse()

	return ServerConfigs{port: *portPtr}
}

// Opens multipart request and returns first file as multipart
func getMultipartFile(req http.Request) (file multipart.File, status int, response string, err error) {
	if err = req.ParseMultipartForm(_24K); nil != err {
		status = http.StatusInternalServerError
		return
	}

	for _, fheaders := range req.MultipartForm.File {
		for _, hdr := range fheaders {

			// open uploaded
			if file, err = hdr.Open(); nil != err {
				status = http.StatusInternalServerError
				response = "Can't open uploading file!"
				return
			}

			return
		}
	}

	status = http.StatusUnprocessableEntity
	response = "Can't find name of file!"
	return
}

// Checks method and write answer to the client
func checkMethodPOSTAndWriteAnswer(w http.ResponseWriter, r *http.Request) error {
	if strings.ToUpper(r.Method) != "POST" {
		writeError(http.StatusMethodNotAllowed, MESSAGE_METHOD_NOT_ALLOWED, w)
		return errors.New(MESSAGE_METHOD_NOT_ALLOWED)
	}

	return nil
}

// Writes status and response to the client
func writeError(status int, response string, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// Writes bytes to the response
func writeFile(bytes []byte, w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=result.jpg")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))

	w.Write(bytes)
}

// Handler for resizing image
func resizeHandler(w http.ResponseWriter, r *http.Request) {
	if nil != checkMethodPOSTAndWriteAnswer(w, r) {
		return
	}

	file, status, response, err := getMultipartFile(*r)
	if nil != err {
		writeError(status, response, w)
		return
	}

	bytes, err := resizing.Resize(file)
	if nil != err {
		writeError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	writeFile(bytes, w, *r)
}

// Handler for creating thumbnail
func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	if nil != checkMethodPOSTAndWriteAnswer(w, r) {
		return
	}

	file, status, response, err := getMultipartFile(*r)
	if nil != err {
		writeError(status, response, w)
		return
	}

	bytes, err := resizing.Thumbnail(file)
	if nil != err {
		writeError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	writeFile(bytes, w, *r)
}

//Entry point of the program
func main() {
	configs := parseArgsToConfigs()

	http.HandleFunc("/resize", resizeHandler)
	http.HandleFunc("/thumbnail", thumbnailHandler)
	http.ListenAndServe(":"+configs.port, nil)
}
