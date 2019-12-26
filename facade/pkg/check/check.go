package check

import (
	"fmt"
	"strings"
)

//CheckFormat struct
type CheckFormat struct {
	Format string
}

//NewCheckFormat ...
func NewCheckFormat(format string) *CheckFormat {
	return &CheckFormat{
		Format: format,
	}
}

//Check ...
func (c *CheckFormat) Check(file string) error {
	format := strings.LastIndex(file, ".")
	if c.Format != file[format+1:] {
		return fmt.Errorf("format not available")
	}
	fmt.Println("format OK")
	return nil
}
