package upload

import "fmt"

// Uploader interface
type Uploader interface {
	Upload(string) error
}

type fileUpload struct {}

// Upload audio file
func (f *fileUpload) Upload(file string) error {
	if file == "" {
		return fmt.Errorf("no file")
	}
	fmt.Println("start uploading")
	fmt.Println("upload done. new file: ", file)
	return nil
}

// NewUploader create new uploader
func NewUploader() Uploader {
	return &fileUpload{}
}
