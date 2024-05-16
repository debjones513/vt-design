package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vt-design/go/src/vtcs/internal/uploader"
)

func failed(s string, err error) bool {
	if err != nil {
		fmt.Printf("Location: %s Error: %s", s, err)
		return true
	}
	return false
}

func main() {

	var eu uploader.ExeUpload
	var url = "http://localhost:8080/upload/api/vtserver_test"
	var testFile = "/Users/debjo/GitHub/vt-design/go/bin/vtserver_test"

	// Input the exe name

	eu.SetExeName(testFile)

	// Convert eu to JSON

	jsonData, err := json.Marshal(eu)
	if failed("json.Marshal", err) {
		log.Fatal(err)
	}

	// Post

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if failed("http.Post", err) {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}
