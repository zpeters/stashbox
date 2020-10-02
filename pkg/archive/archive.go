package archive

import (
	"fmt"
	"io/ioutil"
	"path"
)

func GetArchives(p string) (archives []string, err error) {
	domains, err := ioutil.ReadDir(p)
	if err != nil {
		return archives, err
	}

	for _, domain := range domains {
		domainPath := path.Join(p, domain.Name())
		domainArchives, err := ioutil.ReadDir(domainPath)
		if err != nil {
			return archives, err
		}
		for _, archive := range domainArchives {
			fullArchive := fmt.Sprintf("%s - %s", domain.Name(), archive.Name())
			archives = append(archives, fullArchive)
		}
	}

	return archives, nil
}
