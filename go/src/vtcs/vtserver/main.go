package main

import (
	"fmt"
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

	object_name := r.URL.Path[len("/upload/"):]

	p, err := loadExeOrUrl(object_name)
	if err != nil {
		fmt.Fprintf(w, FmtDefaultErrorHandler)
		log.Fatal(err)
		return
	}

	fmt.Fprintf(w, FmtUploadHandler, p.ExeName, p.ExeSha256, p.ExeBytes)
}

func loadExeOrUrl(object_name string) (*exe_upload.ExeUpload, error) {

	eu, err := exe_upload.Init(object_name)
	return eu, err
}

func main() {

	// Start the web server

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/upload/", uploadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // http://localhost:8080/
}
