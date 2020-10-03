package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"stashbox/pkg/archive"
	"stashbox/pkg/crawler"
	"sync"
)

var crawl bool
var basePath string
var listFlag bool

func main() {
	// get flags
	flag.BoolVar(&crawl, "crawl", false, "crawl and save websites")
	flag.StringVar(&basePath, "b", "./stashDb", "folder to save stash into")
	flag.BoolVar(&listFlag, "list", false, "list saved archives")
	flag.Parse()

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

	if crawl {
		var n int
		var urls []string
		scanner := bufio.NewScanner(os.Stdin)

		// inputs
		fmt.Println("Enter number of urls: ")
		fmt.Scan(&n)
		fmt.Println("Enter the urls: ")
		for i := 0; i < n; i++ {
			scanner.Scan()
			text := scanner.Text()
			if len(text) != 0 {
				urls = append(urls, text)
			} else {
				break
			}
		}

		// spin goroutine for each url 
		var wg sync.WaitGroup
		wg.Add(len(urls))
		for _, url := range urls {
			go crawlUtil(url, &wg)
		}
		wg.Wait()
	} else {
		flag.Usage()
		os.Exit(0)

	}
}

func crawlUtil(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	c, err := crawler.NewCrawler(basePath)
	if err != nil {
		panic(err)
	}

	c.AddUrl(url)
	err = c.Crawl()
	if err != nil {
		fmt.Printf("\nError : %s\n", url)
		fmt.Print(err)
		return
	}

	err = c.Save()
	if err != nil {
		fmt.Printf("\nError : %s\n", url)
		fmt.Print(err)
		return
	}
}
