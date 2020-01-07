package main

import (
	"fmt"

	"github.com/santa-freddy/go_progress/facade/pkg/check"
	"github.com/santa-freddy/go_progress/facade/pkg/convert"
	"github.com/santa-freddy/go_progress/facade/pkg/facade"
	"github.com/santa-freddy/go_progress/facade/pkg/upload"
)

func main() {
	var file = "Tears_Of_The_Dragon.flac"

	facade := facade.NewServiceFacade(
		check.NewCheckFormat("flac"),
		convert.NewConverter("mp3"),
		upload.NewUploader(),
		)
	if err := facade.UploadAudio(file); err != nil {
		fmt.Println("error: ", err.Error())
	}
}
