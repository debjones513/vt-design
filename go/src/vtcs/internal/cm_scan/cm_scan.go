package cm_scan

import (
	"fmt"
	"io"
	"net/http"
)

const CloudMersiveApiKey string = "78e44eaa-3a6d-46f5-a04c-0eea7a33ddf0"
const CloudMersiveScanUrl string = "https://api.cloudmersive.com/virus/scan/cloud-storage/aws-s3/single"

// const CloudMersiveKeysUrl string = "https://portal.cloudmersive.com/keys"

func Scan() {

	url := CloudMersiveScanUrl
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("accessKey", "")
	req.Header.Add("secretKey", "")
	req.Header.Add("bucketRegion", "")
	req.Header.Add("bucketName", "")
	req.Header.Add("keyName", "")
	req.Header.Add("Apikey", CloudMersiveApiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
