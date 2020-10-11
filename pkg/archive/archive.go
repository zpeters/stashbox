package archive

import (
	"fmt"
	"strings"

	"github.com/bmatcuk/doublestar/v2"
)

type archive struct {
	Url   string
	Dates []string
}

// Return a list of archives in the stash
func GetArchives(basePath string) (archives []archive, err error) {
	files, err := doublestar.Glob(fmt.Sprintf("%s/**/*.pdf", basePath))
	if err != nil {
		return archives, err
	}
	return buildArchives(basePath, files), err
}

func buildArchives(path string, files []string) []archive {
	var archives []archive
	var pPage string
	var page string
	dates := make([]string, 0)
	path += "/"
	for _, file := range files {
		pieces := strings.Split(strings.TrimPrefix(file, strings.TrimLeft(path, "./")), "/")
		page = strings.Join(pieces[0:len(pieces)-1], "/")
		date := strings.TrimRight(pieces[len(pieces)-1], ".pdf")

		if page != pPage && pPage != "" {
			a := archive{pPage, dates}
			archives = append(archives, a)
			dates = make([]string, 0)
		}

		dates = append(dates, date)
		pPage = page
	}

	a := archive{page, dates}
	archives = append(archives, a)

	return archives
}
