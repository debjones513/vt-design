// Package cmscanner implements utility routines for using Cloud Mersive to scan a binary
//
// Uploaded binaries are retained forever. The file must be uploaded to S3.
// To upload files, see package [vtcs/internal/uploader]
package cmscanner

import (
	"fmt"
	"net/http"
)

const cloudMersiveAPIKey string = "78e44eaa-3a6d-46f5-a04c-0eea7a33ddf0"
const cloudMersiveScanURL string = "https://api.cloudmersive.com/virus/scan/cloud-storage/aws-s3/single"

// const CloudMersiveKeysUrl string = "https://portal.cloudmersive.com/keys"

func failed(s string, err error) bool {

	if err != nil {
		fmt.Printf("Location: %s Error: %s", s, err)
		return true
	}
	return false
}

// Scan a binary or URL
func Scan() error {

	// https://api.cloudmersive.com/go-client.asp

	url := cloudMersiveScanURL
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if failed("http.NewRequest", err) {
		return err
	}

	req.Header.Add("accessKey", "")
	req.Header.Add("secretKey", "")
	req.Header.Add("bucketRegion", "")
	req.Header.Add("bucketName", "")
	req.Header.Add("keyName", "")
	req.Header.Add("Apikey", cloudMersiveAPIKey)

	res, err := client.Do(req)
	if failed("client.Do", err) {
		return err
	}
	defer res.Body.Close()

	//body, err := io.ReadAll(res.Body)
	//if failed("io.ReadAll", err) {
	//	return err
	//}
	return nil
}
