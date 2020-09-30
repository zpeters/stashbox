package crawler

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"

	"jaytaylor.com/html2text"
)

type Site struct {
	HtmlBody []byte
	TextBody []byte
	Url      string
	Title    string
}

type Crawler struct {
	Urls    []string
	Archive string
	Sites   []Site
}

func NewCrawler(archive string) (Crawler, error) {
	return Crawler{
		Archive: archive,
	}, nil
}

// TODO optimize this
func (c *Crawler) Save() error {
	ensureArchive(c.Archive)
	for _, s := range c.Sites {
		fmt.Printf("Saving %s...\n", s.Url)
		parsed, err := url.Parse(s.Url)
		if err != nil {
			return err
		}
		d := parsed.Host

		// generate a file title
		h := sha1.New()
		io.WriteString(h, s.Url)
		s.Title = fmt.Sprintf("%x", h.Sum(nil))

		// create the domain folder if needed
		err = os.MkdirAll(path.Join(c.Archive, d), 0777)
		if err != nil {
			return err
		}

		// save the html
		htmlFileName := fmt.Sprintf("%s.html", s.Title)
		htmlSavePath := path.Join(c.Archive, d, htmlFileName)
		err = ioutil.WriteFile(htmlSavePath, s.HtmlBody, 0777)
		if err != nil {
			return err
		}

		// save the text
		textFileName := fmt.Sprintf("%s.txt", s.Title)
		textSavePath := path.Join(c.Archive, d, textFileName)
		err = ioutil.WriteFile(textSavePath, s.TextBody, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Crawler) AddUrl(url string) {
	c.Urls = append(c.Urls, url)
}

func (c *Crawler) Crawl() error {
	for _, u := range c.Urls {
		fmt.Printf("Crawling %s...\n", u)

		var site Site

		htmlBody, err := getHtmlBody(u)
		if err != nil {
			return err
		}
		site.HtmlBody = htmlBody

		textBody, err := getTextBody(htmlBody)
		if err != nil {
			return err
		}
		site.TextBody = textBody

		site.Url = u

		c.Sites = append(c.Sites, site)
	}
	return nil
}

func getHtmlBody(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	htmlBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return htmlBody, err
}

func getTextBody(htmlBody []byte) (body []byte, err error) {
	text, err := html2text.FromString(string(htmlBody), html2text.Options{PrettyTables: true})
	if err != nil {
		return body, err
	}

	return []byte(text), nil
}

func ensureArchive(p string) {
	err := os.MkdirAll(p, 0777)
	if err != nil {
		panic(err)
	}
}
