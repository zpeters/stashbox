package crawler

import (
	"testing"
)

func TestDummy(t *testing.T) {

}

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
	c, err := NewCrawler("")
	if err != nil {
		t.Errorf("Unable to create crawler:" + err.Error())
	}

	url := "https://www.github.com"
	err = c.AddUrl(url)
	if err != nil {
		t.Errorf("Test case for url: '" + url + "' failed; it should pass; error:" + err.Error())
	}

	url = "http://www.github.com:8000/"
	err = c.AddUrl(url)
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
	if err != nil {
		t.Errorf("Test case for url: '" + url + "' failed; it should pass; error:" + err.Error())
	}
}
