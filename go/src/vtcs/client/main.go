package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vt-design/go/src/vtcs/internal/exe_upload"
)

func main() {

	var eu exe_upload.ExeUpload
	var url = "http://localhost:8080/upload/api/vtserver_test"
	var test_file = "/Users/debjo/GitHub/vt-design/go/bin/vtserver_test"

	// Input the exe name

	eu.SetExeName(test_file)

	// Convert eu to JSON

	jsonData, err := json.Marshal(eu)
	if err != nil {
		log.Fatal(err)
	}

	// Post

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}
