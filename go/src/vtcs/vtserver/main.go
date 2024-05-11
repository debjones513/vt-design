package main

import (
	"fmt"
	"os"
	"vt-design/go/src/vtcs/internal/exe_upload"
)

func loadExe(name string) *exe_upload.ExeUpload {
	ep := new(exe_upload.ExeUpload)
	ep.ExeName = name + ".txt"
	ep.ExeBytes, _ = os.ReadFile(ep.ExeName)

	return ep
}
func main() {
	eps := loadExe("vtserver_test")
	fmt.Println(eps.ExeName)
	fmt.Println(eps.ExeBytes)
}
