package zget

import (
	"net/http"
	"os"
	"io"
	"log"
)

const (
	OkCode = 0
	BadRequestCode = 1
	BadFileCreationCode = 2
	BadFileCopyCode = 3
)

type fileData struct {
	Path string
	Code int
}

func Download(url string, filePath string) (fileData, error) {
	var err error

	out, err = os.Create(filePath)

	if err != nil {
		return fileData {
			Path: "",
			Code: BadFileCreationCode,
		}, err
	}

	defer out.Close()

	resp, err = http.Get(url)

	if err != nil {
		return fileData {
			Path: "",
			Code: BadRequestCode,
		}, err
	}

	defer resp.Body.Close()

	_, err := io.Copy(out, resp.Body)

	if err != nil {
		return fileData {
			Path: "",
			Code: BadFileCopyCode,
		}, err
	}

	return fileData {
		Path: filePath,
		Code: OkCode,
	}, nil
}
