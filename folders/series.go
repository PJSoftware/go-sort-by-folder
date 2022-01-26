package folders

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

// Series folder data
type Series struct {
	name    string
	count   int
	seasons []Season
}

func (s *Series) scan(parent string) {
	files, err := ioutil.ReadDir(parent + s.name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			season := new(Season)
			season.name = file.Name()
			season.scan(parent + s.name + "/")
			if season.count != 0 {
				s.seasons = append(s.seasons, *season)
				s.count += season.count
			}
		}
	}

}

// Name returns the name of the Series
func (s *Series) Name() string {
	return s.name
}

// Rename does all the heavy lifting
func (s *Series) Rename() {
	fmt.Printf("\nRenaming %d video files in '%s':\n", s.count, s.name)
	vList := s.Sorted()
	for _, video := range vList {
		video.Move()
	}
}

// Sorted returns a list of sorted videos in this series
func (s *Series) Sorted() []Video {
	var rv []Video
	sort.Slice(s.seasons, func(i, j int) bool {
		return s.seasons[i].name < s.seasons[j].name
	})
	for _, season := range s.seasons {
		for _, video := range season.Sorted() {
			video.new = "./" + s.name + "/" + video.new
			rv = append(rv, video)
		}
	}
	return rv
}
