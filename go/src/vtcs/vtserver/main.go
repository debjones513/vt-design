package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"vt-design/go/src/vtcs/internal/exe_upload"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Web server is running %s!", r.URL.Path[1:])
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	object_name := r.URL.Path[len("/upload/"):]
	p, _ := loadExe(object_name)
	fmt.Fprintf(w, "<p style=\"color:green;\">Filename: %s</p><div><p style=\"color:blue;\">Sha256: %s</p></div><div><p style=\"color:black;\">Bytes:</p> <p style=\"color:gray;\">%s</p></div>", p.ExeName, p.ExeSha256, p.ExeBytes)
}

func loadExe(name string) (*exe_upload.ExeUpload, error) {

	ep := new(exe_upload.ExeUpload)

	//
	// Set the file name
	//

	ep.ExeName = "/Users/debjo/GitHub/vt-design/go/bin/" + name

	//
	// Set the file bytes
	//

	data, err := os.ReadFile(ep.ExeName)
	if err != nil {
		log.Fatal(err)
	}

	ep.ExeBytes = make([]byte, len(data))
	n := copy(ep.ExeBytes, data)
	if n != len(data) {
		log.Fatal(err)
	}

	//
	// Set the file hash signature
	//

	f, err := os.Open(ep.ExeName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	ep.ExeSha256 = make([]byte, len(h.Sum(nil)))
	copy(ep.ExeSha256, h.Sum(nil))

	return ep, nil
}
func main() {

	// Start the web server
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/upload/", uploadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // http://localhost:8080/
}
