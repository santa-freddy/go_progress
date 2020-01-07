package facade

import (
	"fmt"
)

type checker interface {
	Check(string) error
}

type converter interface {
	Convert (string) string
}

type uploader interface {
	Upload(string) error
}

type serviceFacade struct {
	check checker
	convert converter
	upload uploader
}

//UploadAudio upload audio file
func (s *serviceFacade) UploadAudio(file string) error {
	if err := s.check.Check(file); err != nil {
		return fmt.Errorf("check err: %s", err.Error())
	}

	audioFile := s.convert.Convert(file)

	if err := s.upload.Upload(audioFile); err != nil {
		return fmt.Errorf("upload err: %s", err.Error())
	}

	return nil
}

//NewServiceFacade create new facade facade
func NewServiceFacade(
	check checker,
	convert converter,
	upload uploader,
	) *serviceFacade {
	return &serviceFacade{
		check:   check,
		convert: convert,
		upload:  upload,
	}
}
