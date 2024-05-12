package main

import (
	"fmt"
	"log"
	"net/http"
	"vt-design/go/src/vtcs/internal/exe_upload"
)

const FmtDefaultHandler string = `
<div> <p style="color:green;"> &nbsp; &nbsp; <b> Web server is running %s! </b> </p> </div>
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
	p, _ := loadExe(object_name)
	//fmt.Fprintf(w, "<p style=\"color:green;\"> &nbsp; &nbsp; <b> Filename </b></p> <p style=\"color:green;\"> &nbsp; &nbsp; %s</p> <div><p style=\"color:blue;\"> &nbsp; &nbsp; <b> Sha256 </b> </p> <p style=\"color:blue;\"> &nbsp; &nbsp; %x</p></div><div><p style=\"color:black;\"> &nbsp; &nbsp; <b> Bytes </b></p> <p style=\"color:gray;\"> &nbsp; &nbsp; %x</p></div>", p.ExeName, p.ExeSha256, p.ExeBytes)
	fmt.Fprintf(w, FmtUploadHandler, p.ExeName, p.ExeSha256, p.ExeBytes)
}

func loadExe(name string) (*exe_upload.ExeUpload, error) {

	eu, err := exe_upload.Init(name)

	return eu, err
}

func main() {

	// Start the web server
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/upload/", uploadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // http://localhost:8080/
}
