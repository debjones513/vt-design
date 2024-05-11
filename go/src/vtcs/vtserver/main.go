package main

import (
	"fmt"
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
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.ExeName, p.ExeBytes)
}

func loadExe(name string) (*exe_upload.ExeUpload, error) {

	data, err := os.ReadFile("/Users/debjo/GitHub/vt-design/go/bin/" + name)
	if err != nil {
		log.Fatal(err)
	}

	ep := new(exe_upload.ExeUpload)
	ep.ExeBytes = make([]byte, len(data))
	n := copy(ep.ExeBytes, data)
	if n != len(data) {
		log.Fatal(err)
	}
	ep.ExeName = name

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
