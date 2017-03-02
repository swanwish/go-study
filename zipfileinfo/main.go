package main

import (
	"archive/zip"
	"os"

	"github.com/swanwish/go-common/logs"
)

func main() {
	file, err := os.Open("/Users/Stephen/Downloads/test.epub")
	if err != nil {
		logs.Errorf("Failed to open file, the error is %v", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		logs.Errorf("Failed to get file info, the error is %v", err)
		return
	}

	zipFileInfoHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		logs.Errorf("Failed to get zip file info header, the error is %v", err)
		return
	}
	logs.Debugf("The coment is %s", zipFileInfoHeader.Comment)
	logs.Debugf("The extra info is %s", string(zipFileInfoHeader.Extra))

	zipFileInfoHeader.Comment = "test comment"
	zipFileInfoHeader.Extra = []byte("passwords")
}
