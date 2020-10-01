package main

import (
	"flag"
	"fmt"
	"os"
	"stashbox/pkg/crawler"
)

// TODO implement functional options
// TODO implement pdf rendering
// TODO figure out a better way to generate a file name
// TODO how do we handle revisions
// TODO tests

var u string
var basePath string

func main() {
	// get flags
	flag.StringVar(&u, "url", "", "url to download")
	flag.StringVar(&basePath, "b", "./stashDb", "folder to save stash into")
	flag.Parse()

	if u == "" {
		fmt.Println("Please supply a url.  See -h for more info")
		os.Exit(1)
	}

	c, err := crawler.NewCrawler(basePath)
	if err != nil {
		panic(err)
	}

	c.AddUrl(u)
	err = c.Crawl()
	if err != nil {
		panic(err)
	}

	err = c.Save()
	if err != nil {
		panic(err)
	}
}
