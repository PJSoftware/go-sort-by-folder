package folders

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

// Season folder data
type Season struct {
	name  string
	count int
	disks []Disk
}

func (s *Season) scan(parent string) {
	files, err := ioutil.ReadDir(parent + s.name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			disk := new(Disk)
			disk.name = file.Name()
			disk.scan(parent+s.name+"/", s.name)
			if disk.count != 0 {
				s.disks = append(s.disks, *disk)
				s.count += disk.count
			}
		}
	}

}

// Sorted returns a list of sorted videos in this season
func (s *Season) Sorted() []Video {
	var rv []Video
	sort.Slice(s.disks, func(i, j int) bool {
		return s.disks[i].name < s.disks[j].name
	})
	ep := 1
	for _, disk := range s.disks {
		for _, video := range disk.Sorted() {
			video.new = s.name + "E" + fmt.Sprintf("%02d", ep) + "." + video.ext
			ep++
			rv = append(rv, video)
		}
	}
	return rv
}
