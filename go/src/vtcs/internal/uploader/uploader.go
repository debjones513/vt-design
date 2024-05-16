// Package uploader implements utility routines for uploading binaries to S3.
//
// Uploaded binaries are retained forever.
package uploader

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// ExeUpload identifies a binary by name and signature. It includes the binary's
// bytes as a byte array.
// An empty string for ExeName is an uninitialized structure.
// If the stricture is initialized, ExeBytes and ExeSha256 are assumed to be of
// length greater than 0.
type ExeUpload struct {
	ExeName   string // The executable's filename
	ExeBytes  []byte // The executable's bytes
	ExeSha256 []byte // The SHA256 for the executable's bytes
}

// https://aws.github.io/aws-sdk-go-v2/docs/making-requests/
// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2
// https://github.com/aws/aws-sdk-go-v2/blob/example/service/s3/listObjects/v0.2.9/example/service/s3/listObjects/listObjects.go
// https://docs.aws.amazon.com/AmazonS3/latest/userguide/example_s3_Scenario_PresignedUrl_section.html
// https://ronen-niv.medium.com/aws-s3-handling-presigned-urls-2718ab247d57

/*
https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/s3#PutObjectInput


func putPresignURL(cfg aws.Config) string {

s3client := s3.NewFromConfig(cfg)
 presignClient := s3.NewPresignClient(s3client)

 presignedUrl, err := presignClient.PresignPutObject(context.Background(),
  &s3.PutObjectInput{
   Bucket: aws.String(bucketName),
   Key:    aws.String(objectName),
  },
  s3.WithPresignExpires(time.Minute*15))
 if err != nil {
  log.Fatal(err)
 }

 return presignedUrl.URL
}


func uploadFile(filePath string, url string) error {

 file, err := os.Open(filePath)
 if err != nil {
  return err
 }
 defer file.Close()

 buffer := bytes.NewBuffer(nil)
 if _, err := io.Copy(buffer, file); err != nil {
  return err
 }

 request, err := http.NewRequest(http.MethodPut, url, buffer)
 if err != nil {
  return err
 }

 request.Header.Set("Content-Type", "multipart/form-data")
 client := &http.Client{}
 _, err = client.Do(request)
 return err
}

*/

func failed(s string, err error) bool {
	if err != nil {
		fmt.Printf("Location: %s Error: %s", s, err)
		return true
	}
	return false
}

// SetExeName copies the executable name into a string.
func (eu *ExeUpload) SetExeName(name string) error {

	// Set the file name

	eu.ExeName = name

	return nil
}

// SetExeBytes copies the executable's bytes into a byte array.
func (eu *ExeUpload) SetExeBytes() error {

	// Set the file bytes - TODO for a big file, this will consume memory

	data, err := os.ReadFile(eu.ExeName)
	if failed("os.ReadFile", err) {
		return err
	}

	eu.ExeBytes = make([]byte, len(data))
	n := copy(eu.ExeBytes, data)
	if n != len(data) {
		fmt.Printf("Copy", err)
		return err
	}

	return nil
}

// SetExeSha256 generates a SHA256 (signature) from the executable's bytes. Copies the executable's signature into a byte array.
func (eu *ExeUpload) SetExeSha256() error {

	// Set the file hash signature

	f, err := os.Open(eu.ExeName)
	if failed("os.Open", err) {
		return err
	}
	defer f.Close()

	// Hash the file bytes - batches the file bytes into the hash to minimize memory usage

	h := sha256.New()
	_, err = io.Copy(h, f)
	if failed("io.Copy", err) {
		return err
	}

	// Copy hash to struct

	eu.ExeSha256 = make([]byte, len(h.Sum(nil)))
	n := copy(eu.ExeSha256, h.Sum(nil))
	if !(n > 0) {
		fmt.Printf("copy", io.EOF)
		return io.EOF
	}

	return nil
}

// Initialize initializes an ExeUpload object.
func Initialize(name string) (*ExeUpload, error) {

	var err error

	eu := new(ExeUpload)

	err = eu.SetExeName(name)
	if failed("SetExeName", err) {
		return eu, err
	}

	err = eu.SetExeSha256()
	if failed("SetExeSha256", err) {
		return eu, err
	}

	err = eu.SetExeBytes()
	if failed("SetExeBytes", err) {
		return eu, err
	}

	return eu, nil
}
