package convert

import (
	"fmt"
	"strings"
)

// Converter interface
type Converter interface {
	Convert(string) string
}

type audio struct {
	ResultFormat string
}

// Convert audio file
func (a *audio) Convert(name string) string {
	fmt.Println("start converting")
	newName := strings.Replace(name, name[strings.LastIndex(name, ".")+1:], a.ResultFormat, 1)
	return newName
}

// NewConverter create new converter
func NewConverter(format string) Converter {
	return &audio{
		ResultFormat: format,
	}
}
