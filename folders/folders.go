package folders

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// FolderContents is our master 'folders' data structure
type FolderContents struct {
	series map[string]*Series
}

// Init reads current folder tree, builds FolderContents
func Init() *FolderContents {
	fc := newFolderContents()
	fc.scan(".")
	return fc
}

// ChooseSeries asks user to select from those with sortable files
func (fc *FolderContents) ChooseSeries() *Series {
	if len(fc.series) == 0 {
		fmt.Println("No series found requiring sorting!")
		return nil
	}

	valid := ""
	var keys []string
	for key := range fc.series {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println("The following series are available to be sorted:")
	for _, key := range keys {
		sOpt := fmt.Sprintf("  (%s) %s: ", key, fc.series[key].name)
		for idx, season := range fc.series[key].seasons {
			if idx > 0 {
				sOpt += "; "
			}
			sOpt += fmt.Sprintf("%s (%d eps)", season.name, season.count)
		}
		fmt.Println(sOpt)
		valid += key
	}
	choice := getChoice("Select series", valid)
	if choice == "" {
		return nil
	}
	return fc.series[choice]
}

func newFolderContents() *FolderContents {
	fc := new(FolderContents)
	fc.series = make(map[string]*Series)
	return fc
}

func (fc *FolderContents) scan(parent string) {
	files, err := ioutil.ReadDir(parent)
	if err != nil {
		log.Fatal(err)
	}

	key := 'a'
	for _, file := range files {
		if file.IsDir() {
			series := new(Series)
			series.name = file.Name()
			series.scan(parent + "/")
			if series.count > 0 {
				fc.series[string(key)] = series
				key++
			}
		}
	}

}

func getChoice(prompt string, valid string) string {
	reader := bufio.NewReader(os.Stdin)
	if valid != "" {
		prompt = prompt + " [" + valid + "]"
	}

	validPattern := "^[" + valid + "]$"
	var isValid bool
	for !isValid {
		fmt.Print(prompt + ": ")
		r, _ := reader.ReadString('\n')
		r = strings.Replace(r, "\r", "", -1)
		r = strings.Replace(r, "\n", "", -1)
		if r == "" {
			return r
		}
		if ok, _ := regexp.MatchString(validPattern, r); ok {
			return r
		}
	}
	return ""
}
