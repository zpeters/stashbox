package archive

import (
	"fmt"
	"strings"

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
	var archives = []Archive{}
	var pPage string
	var page string
	dates := make([]string, 0)
	path += "/"
	for _, file := range files {
		pieces := strings.Split(strings.TrimPrefix(file, strings.TrimLeft(path, "./")), "/")
		page = strings.Join(pieces[0:len(pieces)-1], "/")
		date := strings.TrimRight(pieces[len(pieces)-1], ".pdf")

		if page != pPage && pPage != "" {
			return archives, fmt.Errorf("Error building archive")
		}

		dates = append(dates, date)
		pPage = page
	}

	a := Archive{page, dates}
	archives = append(archives, a)

	return archives, nil
}
