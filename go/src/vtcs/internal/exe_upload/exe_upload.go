package exe_upload

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

type ExeUpload struct {
	ExeName   string
	ExeBytes  []byte
	ExeSha256 []byte
}

func (eu *ExeUpload) SetExeName(name string) error {

	// Set the file name

	eu.ExeName = "/Users/debjo/GitHub/vt-design/go/bin/" + name

	return nil
}

func (eu *ExeUpload) SetExeBytes() error {

	// Set the file bytes - TODO for a big file, this will consume memory

	data, err := os.ReadFile(eu.ExeName)
	if err != nil {
		log.Fatal(err)
		return err
	}

	eu.ExeBytes = make([]byte, len(data))
	n := copy(eu.ExeBytes, data)
	if n != len(data) {
		log.Fatal(err)
		return err
	}

	return nil
}

func (eu *ExeUpload) SetExeSha256() error {

	// Set the file hash signature

	f, err := os.Open(eu.ExeName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	// Hash the file bytes - batches the file bytes into the hash to minimize memory usage

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Copy hash to struct

	eu.ExeSha256 = make([]byte, len(h.Sum(nil)))
	n := copy(eu.ExeSha256, h.Sum(nil))
	if !(n > 0) {
		log.Fatal(io.EOF)
		return io.EOF
	}

	return nil
}

func Init(name string) (*ExeUpload, error) {

	var err error

	eu := new(ExeUpload)

	err = eu.SetExeName(name)
	if err != nil {
		log.Fatal(err)
		return eu, err
	}

	err = eu.SetExeBytes()
	if err != nil {
		log.Fatal(err)
		return eu, err
	}

	err = eu.SetExeSha256()
	if err != nil {
		log.Fatal(err)
		return eu, err
	}

	return eu, nil
}
