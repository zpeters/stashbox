package crawler

import (
	"strconv"
	"testing"
	"time"
)

func TestGetHtmlTitle(t *testing.T) {
	const url = "https://github.com/zpeters/stashbox"
	const want = "GitHub - zpeters/stashbox: Your personal Internet Archive"

	body, err := getHtmlBody(url)
	if err != nil {
		t.Errorf(err.Error())
	}

	title, err := getHtmlTitle(body)
	if err != nil {
		t.Errorf(err.Error())
	}

	if title != want {
		t.Errorf("Wrong title found. Want: %s, Got : %s", want, title)
	}
}

func TestAddUrl(t *testing.T) {
	count := 0
	c, err := NewCrawler("")
	if err != nil {
		t.Errorf("Unable to create crawler:" + err.Error())
	}

	url := "https://www.github.com"
	err = c.AddUrl(url)
	count++
	if err != nil {
		t.Errorf("Test case for url: '" + url + "' failed; it should pass; error:" + err.Error())
	}

	url = "http://www.github.com:8000/"
	err = c.AddUrl(url)
	count++
	if err != nil {
		t.Errorf("Test case for url: '" + url + "' failed; it should pass; error:" + err.Error())
	}

	url = "httpd://www.github.com"
	err = c.AddUrl(url)
	if err == nil {
		t.Errorf("Test case for url: '" + url + "' failed; it should fail, but we got no error")
	}

	url = "https:///www.github.com"
	err = c.AddUrl(url)
	if err == nil {
		t.Errorf("Test case for url: '" + url + "' failed; it should fail, but we got no error")
	}

	url = "//www.github.com:8000"
	err = c.AddUrl(url)
	if err == nil {
		t.Errorf("Test case for url: '" + url + "' failed; it should fail, but we got no error")
	}

	url = "www.github.com:8000/"
	err = c.AddUrl(url)
	count++
	if err != nil {
		t.Errorf("Test case for url: '" + url + "' failed; it should pass; error:" + err.Error())
	}

	if len(c.Urls) != count {
		t.Errorf("There should be " + strconv.Itoa(count) + " entries, found:" + strconv.Itoa(len(c.Urls)))
	}
}

func TestBuildPath(t *testing.T) {
	p, e := buildPath("./StashDB", "http://www.google.com/a/test.html")
	if e != nil {
		t.Error(e)
	}
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
