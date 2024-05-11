package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vt-design/go/src/vtcs/internal/exe_upload"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Web server is running %s!", r.URL.Path[1:])
}

func testLoadExe(name string) *exe_upload.ExeUpload {
	ep := new(exe_upload.ExeUpload)
	ep.ExeName = name + ".txt"
	ep.ExeBytes, _ = os.ReadFile(ep.ExeName)

	return ep
}
func main() {
	eps := testLoadExe("vtserver_test")
	fmt.Println(eps.ExeName)
	fmt.Println(eps.ExeBytes)

	// Start the web server
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
