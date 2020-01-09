package facade

import (
	"testing"

	"github.com/santa-freddy/go_progress/facade/pkg/check"
	"github.com/santa-freddy/go_progress/facade/pkg/convert"
	"github.com/santa-freddy/go_progress/facade/pkg/upload"
)

func TestUploadAudio(t *testing.T) {

	uploadAudioTest := []struct {
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
		{
			"upload_3",
			"file.wav",
			true,
		},
	}

	uploader := NewServiceFacade(
		check.NewCheckFormat("flac"),
		convert.NewConverter("mp3"),
		upload.NewUploader(),
		)

	for _, tt := range uploadAudioTest {
		t.Run(tt.name, func(t *testing.T) {
			if err := uploader.UploadAudio(tt.args); (err != nil) != tt.wantErr {
				t.Errorf(
					"UploadAudio() error = %v, wantErr = %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
