package main

import (
	"flag"
	"fmt"
	"os"
	"stashbox/pkg/archive"
	"stashbox/pkg/crawler"
)

var u string
var basePath string
var listFlag bool

func main() {
	// get flags
	flag.StringVar(&u, "url", "", "url to download")
	flag.StringVar(&basePath, "b", "./stashDb", "folder to save stash into")
	flag.BoolVar(&listFlag, "list", false, "list saved archives")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	if listFlag {
		archives, err := archive.GetArchives(basePath)
		if err != nil {
			panic(err)
		}
		if len(archives) == 0 {
			fmt.Println("No archives found...")
		} else {
			fmt.Println("Archive listing...")
			for n, a := range archives {
				fmt.Printf("%d. %s\n", n+1, a)
			}
		}
		os.Exit(0)
	}

	if u != "" {
		c, err := crawler.NewCrawler(basePath)
		if err != nil {
			panic(err)
		}

		err = c.AddUrl(u)
		if err != nil {
			panic(err)
		}

		err = c.Crawl()
		if err != nil {
			panic(err)
		}

		err = c.Save()
		if err != nil {
			panic(err)
		}
	}
}
