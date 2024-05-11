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
	fmt.Fprintf(w, "<h1>Filename: %s</h1><div><p style=\"color:blue;\">Sha256: %s</p></div><div>Bytes: %s</div>", p.ExeName, p.ExeSha256, p.ExeBytes)
}

func loadExe(name string) (*exe_upload.ExeUpload, error) {

	ep := new(exe_upload.ExeUpload)

	// Set the file name
	ep.ExeName = "/Users/debjo/GitHub/vt-design/go/bin/" + name

	// Get the file bytes
	data, err := os.ReadFile(ep.ExeName)
	if err != nil {
		log.Fatal(err)
	}

	ep.ExeBytes = make([]byte, len(data))
	n := copy(ep.ExeBytes, data)
	if n != len(data) {
		log.Fatal(err)
	}

	// Get the file hash signature
	f, err := os.Open(ep.ExeName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", h.Sum(nil))

	ep.ExeSha256 = make([]byte, len(h.Sum(nil)))
	copy(ep.ExeSha256, h.Sum(nil))

	return ep, nil
}
func main() {
	//eps := loadExe("vtserver_test")
	//fmt.Println(eps.ExeName)
	//fmt.Println(eps.ExeBytes)

	// Start the web server
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/upload/", uploadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // http://localhost:8080/
}
