package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"vt-design/go/src/vtcs/internal/exe_upload"
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

	// Unmarshal
	var req exe_upload.ExeUpload
	err = json.Unmarshal(b, &req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	eu, err := loadExeOrUrl(req.ExeName)
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

func loadExeOrUrl(object_name string) (*exe_upload.ExeUpload, error) {

	eu, err := exe_upload.Init(object_name)
	if err != nil {
		fmt.Println(err)
	}
	return eu, err
}

func main() {

	// Start the web server

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/upload/api/", uploadApiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // http://localhost:8080/
}
