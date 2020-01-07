package check

import (
	"fmt"
	"strings"
)

type checkFormat struct {
	Format string
}

//Checker interface
type Checker interface {
	Check(string) error
}

//Check file format
func (c *checkFormat) Check(file string) error {
	format := strings.LastIndex(file, ".")
	if c.Format != file[format+1:] {
		return fmt.Errorf("format not available")
	}
	fmt.Println("format OK")
	return nil
}

// NewCheckFormat create new checker format
func NewCheckFormat(format string) Checker {
	return &checkFormat{
		Format: format,
	}
}
