package archive

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bmatcuk/doublestar/v2"
)

// Archive is a url with one or more dates (ie recordings of the archive)
type Archive struct {
	URL   string
	Dates []string
}

// GetArchives returns a list of archives in the stash
func GetArchives(basePath string) (archives []Archive, err error) {
	files, err := doublestar.Glob(fmt.Sprintf("%s/**/*.pdf", basePath))
	if err != nil {
		return archives, fmt.Errorf("Error getting archives: %w", err)
	}
	if len(files) == 0 {
		return archives, fmt.Errorf("no archives found in %s", basePath)
	}

	archives, err = buildArchives(basePath, files)
	if err != nil {
		return archives, err
	}

	return archives, nil
}

func buildArchives(path string, files []string) ([]Archive, error) {
	var archive *Archive
	var dates []string
	var pPage string
	var page string

	archives := []Archive{}
	setOfArchives := make(map[*Archive]bool)
	path += "/"

	for _, file := range files {
		pieces := strings.Split(strings.TrimPrefix(file, strings.TrimLeft(path, "./")), "/")
		page = strings.Join(pieces[0:len(pieces)-1], "/")
		date := strings.TrimRight(pieces[len(pieces)-1], ".pdf")

		if _, err := time.Parse("2006-02-01T15:04:05", date); err != nil {
			return archives, fmt.Errorf("Error building archive")
		}
		if page != pPage {
			dates = make([]string, 0)
			archive = &Archive{URL: page}
		}

		dates = append(dates, date)
		archive.Dates = dates

		if _, present := setOfArchives[archive]; !present {
			setOfArchives[archive] = true
		}

		pPage = page

	}

	for key := range setOfArchives {
		archives = append(archives, *key)
	}

	// Sort the archives because it is not sorted when using a map.
	// Important for "open" command to make sure it is in the same
	// order every time getArchives is called
	sort.Slice(archives, func(i, j int) bool {
		return archives[i].URL < archives[j].URL
	})

	return archives, nil
}
