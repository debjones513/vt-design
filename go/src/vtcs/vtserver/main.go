/*
Vtserver runs a virus scan on a binary.
It stores the binary in S3, and uses third party scanners.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"vt-design/go/src/vtcs/internal/uploader"
)

// FmtDefaultHandler is the HTML for default web server handler
const FmtDefaultHandler string = `
<div> <p style="color:green;"> &nbsp; &nbsp; <b> Web server is running %s! </b> </p> </div>
<div> <p style="color:green;"> &nbsp; &nbsp; To test: http://localhost:8080/upload/vtserver_test </p> </div>
`

// FmtDefaultErrorHandler is the HTML for default web server error handler
const FmtDefaultErrorHandler string = `
<div> <p style="color:red;"> &nbsp; &nbsp; <b> ERROR! </b> </p> </div>
`

// FmtUploadHandler is the HTML for web server upload handler
const FmtUploadHandler string = `
<div> <p style="color:green;"> &nbsp; &nbsp; <b> Filename </b> </p>   <p style="color:green;"> &nbsp; &nbsp; %s </p> </div>
<div> <p style="color:blue;">  &nbsp; &nbsp; <b> Sha256   </b> </p>   <p style="color:blue;">  &nbsp; &nbsp; %x </p> </div>
<div> <p style="color:black;"> &nbsp; &nbsp; <b> Bytes    </b> </p>   <p style="color:gray;">  &nbsp; &nbsp; %x </p> </div>
`

func failed(s string, err error) bool {
	if err != nil {
		fmt.Printf("Location: %s Error: %s", s, err)
		return true
	}
	return false
}

// Default web server handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, FmtDefaultHandler, r.URL.Path[1:])
}

// Web server upload handler
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	objectName := "/Users/debjo/GitHub/vt-design/go/bin/" + r.URL.Path[len("/upload/"):]

	eu, err := loadExeOrURL(objectName)
	if err != nil {
		fmt.Fprintf(w, FmtDefaultErrorHandler)
		return
	}

	fmt.Fprintf(w, FmtUploadHandler, eu.ExeName, eu.ExeSha256, eu.ExeBytes)
}

// API upload handler
func uploadAPIHandler(w http.ResponseWriter, r *http.Request) {

	// Read body

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal to get the binary name

	var resp uploader.ExeUpload
	err = json.Unmarshal(b, &resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Load the response
	eu, err := loadExeOrURL(resp.ExeName)
	if err != nil {
		fmt.Fprintf(w, FmtDefaultErrorHandler)
		return
	}

	fmt.Fprintf(w, FmtUploadHandler, eu.ExeName, eu.ExeSha256, eu.ExeBytes)

	/* Marshall the struct into a response
	output, err := json.Marshal(eu)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
	*/
}

func loadExeOrURL(objectName string) (*uploader.ExeUpload, error) {

	eu, err := uploader.Initialize(objectName)
	if failed("uploader.Initialize", err) {
		return eu, err
	}

	//cmscanner.Scan()
	return eu, err
}

func main() {

	// Start the web server

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/upload/api/", uploadAPIHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // http://localhost:8080/
}
