package convert

import (
	"fmt"
	"strings"
)

//Audio struct
type Audio struct {
	ResultFormat string
}

//NewConvert ...
func NewConvert(format string) *Audio {
	return &Audio{
		ResultFormat: format,
	}
}

//Convert ...
func (a *Audio) Convert(name string) string {
	fmt.Println("start converting")
	newName := strings.Replace(name, name[strings.LastIndex(name, ".")+1:], a.ResultFormat, 1)
	return newName
}
