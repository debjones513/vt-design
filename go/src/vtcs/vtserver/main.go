/*
Vtserver runs a virus scan on a binary.
It stores the binary in S3, and uses third party scanners.

The Usage statement below shows how a cmd line app comment should be formatted.
Usage:

	my_cmd line [flags] [path ...]

The flags are:

	-d
	    Do not ...
	-w
	    If ...

Replace the text above with vtserver info.
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

const FmtDefaultHandler string = `
<div> <p style="color:green;"> &nbsp; &nbsp; <b> Web server is running %s! </b> </p> </div>
<div> <p style="color:green;"> &nbsp; &nbsp; To test: http://localhost:8080/upload/vtserver_test </p> </div>
`

const FmtDefaultErrorHandler string = `
<div> <p style="color:red;"> &nbsp; &nbsp; <b> ERROR! </b> </p> </div>
`

const FmtUploadHandler string = `
<div> <p style="color:green;"> &nbsp; &nbsp; <b> Filename </b> </p>   <p style="color:green;"> &nbsp; &nbsp; %s </p> </div>
<div> <p style="color:blue;">  &nbsp; &nbsp; <b> Sha256   </b> </p>   <p style="color:blue;">  &nbsp; &nbsp; %x </p> </div>
<div> <p style="color:black;"> &nbsp; &nbsp; <b> Bytes    </b> </p>   <p style="color:gray;">  &nbsp; &nbsp; %x </p> </div>
`

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, FmtDefaultHandler, r.URL.Path[1:])
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	object_name := "/Users/debjo/GitHub/vt-design/go/bin/" + r.URL.Path[len("/upload/"):]

	eu, err := loadExeOrUrl(object_name)
	if err != nil {
		fmt.Fprintf(w, FmtDefaultErrorHandler)
		return
	}

	fmt.Fprintf(w, FmtUploadHandler, eu.ExeName, eu.ExeSha256, eu.ExeBytes)
}

func uploadApiHandler(w http.ResponseWriter, r *http.Request) {

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
	eu, err := loadExeOrUrl(resp.ExeName)
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

func loadExeOrUrl(object_name string) (*uploader.ExeUpload, error) {

	eu, err := uploader.Initialize(object_name)
	if err != nil {
		fmt.Println(err)
	}

	//cmscanner.Scan()
	return eu, err
}

func main() {

	// Start the web server

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/upload/api/", uploadApiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // http://localhost:8080/
}
