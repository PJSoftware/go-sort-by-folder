package folders

import (
	"fmt"
	"os"
	"regexp"
)

// FileRegexp is used to detect/parse filenames
type FileRegexp struct {
	pattern string
	re      *regexp.Regexp
}

// NewFileRegexp is our constructor
func NewFileRegexp(pattern string) *FileRegexp {
	fr := new(FileRegexp)
	fr.pattern = pattern
	re, err := regexp.Compile(fr.pattern)
	if err != nil {
		fmt.Printf("NewFileRegexp: Error compiling '%s'\n", pattern)
		os.Exit(1)
	}
	fr.re = re
	return fr
}

// Match compares filename with pattern
func (fr *FileRegexp) Match(name string) bool {
	ok, _ := regexp.MatchString(fr.pattern, name)
	return ok
}

// Extract returns matches from filename
func (fr *FileRegexp) Extract(name string) []string {
	rm := fr.re.FindAllStringSubmatch(name, -1)
	return rm[0][1:]
}
