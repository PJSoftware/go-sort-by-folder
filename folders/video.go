package folders

import (
	"fmt"
	"os"
)

// Video tracks details relevant to sorting
type Video struct {
	season string
	disk   string
	track  int
	name   string
	ext    string
	new    string
}

// Move video from name to new
func (v *Video) Move() {
	fmt.Printf("%s -> %s", v.name, v.new)
	err := os.Rename(v.name, v.new)
	if err != nil {
		fmt.Print(": move failed!")
	}
	fmt.Println()
}
