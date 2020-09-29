package main

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// TODO cleanup
// TODO text extraction isn't quite there
// TODO status display,etc
// TODO some sort of hashing or date comparisons.../?
// TODO figure out how to save a new copy every time...symlink latest hash..?  basically  we want to do "append only"
// TODO  tests
// TODO save rendered pdf

var u string
var basePath string

func main() {
	// get flags
	flag.StringVar(&u, "url", "", "url to download")
	flag.StringVar(&basePath, "b", "/users/zachpeters/Downloads/stashDb", "folder to save stash into")
	flag.Parse()

	if u == "" {
		fmt.Println("Please supply a url.  See -h for more info")
		os.Exit(1)
	}

	savePath := path.Join(basePath, getSavePath(u))
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read html
	htmlBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// make dirs if needed
	// BUG why is this needed???
	err = os.MkdirAll(savePath, 0777)
	if err != nil {
		panic(err)
	}

	// get the title
	title, err := getTitle(htmlBody)
	if err != nil {
		panic(err)
	}

	// write to file
	err = ioutil.WriteFile(path.Join(savePath, fmt.Sprintf("%s.html", title)), htmlBody, 0777)
	if err != nil {
		panic(err)
	}

	// get text
	textBody, err := getTextBody(htmlBody)
	if err != nil {
		panic(err)
	}
	// write to file
	err = ioutil.WriteFile(path.Join(savePath, fmt.Sprintf("%s.txt", title)), []byte(textBody), 0777)
	if err != nil {
		panic(err)
	}
}

func getTextBody(htmlBody []byte) (string, error) {
	p := strings.NewReader(string(htmlBody))
	doc, err := goquery.NewDocumentFromReader(p)
	if err != nil {
		return "", err
	}
	doc.Find("script").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})
	return doc.Text(), nil
}

func getTitle(htmlBody []byte) (string, error) {
	doc, err := html.Parse(bytes.NewReader(htmlBody))
	if err != nil {
		panic("Fail to parse html")
	}
	title, ok := traverseDoc(doc)
	if !ok {
		return "", errors.New("No title found, malformed HTML")
	}
	return title, nil
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverseDoc(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverseDoc(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func getSavePath(u string) string {
	// break url down
	parsedUrl, err := url.Parse(u)
	if err != nil {
		panic(err)
	}

	domain := parsedUrl.Host
	pathPart := parsedUrl.Path
	h := sha1.New()
	h.Write([]byte(pathPart))
	bs := h.Sum(nil)

	p := path.Join(domain, fmt.Sprintf("%x", bs))
	return p
}
