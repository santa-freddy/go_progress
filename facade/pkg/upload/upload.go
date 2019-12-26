package upload

import "fmt"

//FileUpload struct
type FileUpload struct {}

//NewUploader ...
func NewUploader() *FileUpload {
	return &FileUpload{}
}

//Upload ...
func (f *FileUpload) Upload(file string) error {
	fmt.Println("start uploading")
	defer fmt.Println("upload done. new file: ", file)
	return nil
}
