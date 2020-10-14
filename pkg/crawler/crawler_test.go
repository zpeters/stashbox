package crawler

import (
	"strconv"
	"testing"
	"time"
)

func TestGetHtmlTitle(t *testing.T) {
	const url = "https://github.com/zpeters/stashbox"
	const want = "GitHub - zpeters/stashbox: Your personal Internet Archive"

	body, err := getHTMLBody(url)
	handleErr(t, err)
	title, err := getHTMLTitle(body)
	handleErr(t, err)
	if title != want {
		t.Errorf("Wrong title found. Want: %s, Got : %s", want, title)
	}
}

func TestAddUrl(t *testing.T) {
	count := 6
	c, err := NewCrawler("")
	if err != nil {
		t.Errorf("Unable to create crawler:" + err.Error())
	}

	for i := 1; i <= count; i++ {
		url := "https://www.github.com" + strconv.Itoa(i)
		err = c.AddURL(url)
		if err != nil {
			t.Errorf("Test case for url: '" + url + "' failed; it should pass; error:" + err.Error())
		}
	}
	if len(c.Urls) != count {
		t.Errorf("There should be " + strconv.Itoa(count) + " entries, found:" + strconv.Itoa(len(c.Urls)))
	}
}

func TestBuildPath(t *testing.T) {
	p, err := buildPath("./StashDB", "http://www.google.com/a/test.html")
	handleErr(t, err)
	expected := "StashDB/www.google.com/a/test.html"
	if p != expected {
		t.Errorf("expected: %s actual: %s", expected, p)
	}
}

func TestDateTimeFileName(t *testing.T) {
	actual := dateTimeFileName()
	expected := time.Now().Format("2006-02-01T15:04:05")
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}

}

func TestCrawl(t *testing.T) {
	var crawlSites = []string{"https://www.google.com", "https://github.com/zpeters/stashbox"}

	c, err := NewCrawler("")
	if err != nil {
		t.Errorf("Unable to create crawler:" + err.Error())
	}

	for _, s := range crawlSites {
		err = c.AddURL(s)
		handleErr(t, err)
	}
	err = c.Crawl()
	handleErr(t, err)

	if len(crawlSites) != len(c.Sites) {
		t.Errorf("got %v sites expected %v sites", len(crawlSites), len(c.Sites))
	}

}

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err.Error())
	}
}
