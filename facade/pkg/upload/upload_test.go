package upload

import (
	"testing"
)

func TestUpload(t *testing.T) {

	uploadTest := []struct {
		name    string
		args    string
		wantErr bool
	 }{
	 	{
	 		"upload_1",
	 		"file.flac",
	 		false,
		},
		{
			"upload_2",
			"",
			true,
		},
	}

	uploader := NewUploader()

	for _, tt := range uploadTest {
		t.Run(tt.name, func(t *testing.T) {
			if err := uploader.Upload(tt.args); (err != nil) != tt.wantErr {
				t.Errorf(
					"Upload() error = %v, wantErr = %v",
					err,
					tt.wantErr,
					)
			}
		})
	}
}