package crawler

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/PuerkitoBio/goquery"
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

// Error Messages
var (
	errNoTitleInHtml = errors.New("No title tag in HTML response")
)

func NewCrawler(archive string) (Crawler, error) {
	return Crawler{
		Archive: archive,
	}, nil
}

func (c *Crawler) Save() error {
	ensureArchive(c.Archive)
	for _, s := range c.Sites {
		fmt.Printf("Saving %s...\n", s.Url)
		parsed, err := url.Parse(s.Url)
		if err != nil {
			return err
		}
		d := parsed.Host

		// get current time
		t := time.Now()
		dateTime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())

		// create the sub folder inside domain folder date-time as name
		domainSubPath := path.Join(c.Archive, d, dateTime)
		err = os.MkdirAll(domainSubPath, 0700)
		if err != nil {
			return err
		}

		// save the html
		htmlFileName := fmt.Sprintf("%s.html", s.Title)
		htmlSavePath := path.Join(domainSubPath, htmlFileName)
		err = ioutil.WriteFile(htmlSavePath, s.HtmlBody, 0600)
		if err != nil {
			return err
		}

		// save the text
		textFileName := fmt.Sprintf("%s.txt", s.Title)
		textSavePath := path.Join(domainSubPath, textFileName)
		err = ioutil.WriteFile(textSavePath, s.TextBody, 0600)
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

		resp, err := getHtmlResponse(u)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		title, err := getHtmlTitle(resp)

		// if there is no title, do old way of creating hash
		if err == errNoTitleInHtml {
			h := sha256.New()
			_, err = io.WriteString(h, u)
			if err != nil {
				return err
			}
			title = fmt.Sprintf("%x", h.Sum(nil))
		} else if err != nil {
			return err
		}
		site.Title = title

		htmlBody, err := getHtmlBody(resp)
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

func getHtmlResponse(url string) (resp *http.Response, err error) {
	// #nosec - gosec will detect this as a G107 error
	// the point of this function *is* to accept a variable URL
	resp, err = http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func getHtmlTitle(resp *http.Response) (title string, err error) {
	// HTML DOM Document
	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		return "", err
	}
	titleTag := doc.Find("title").First()

	if titleTag.Size() == 0 {
		return "", errNoTitleInHtml
	}

	return titleTag.Text(), nil
}

func getHtmlBody(resp *http.Response) (body []byte, err error) {
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
	err := os.MkdirAll(p, 0700)
	if err != nil {
		panic(err)
	}
}
