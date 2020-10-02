package crawler

import (
	"net/http"
	"testing"
)

func TestDummy(t *testing.T) {

}

func TestGetHtmlTitle(t *testing.T) {
	const url = "https://github.com/zpeters/stashbox"
	const want = "GitHub - zpeters/stashbox: Your personal Internet Archive"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf(err.Error())
	}
	title, err := getHtmlTitle(resp)
	if err != nil {
		t.Errorf(err.Error())
	}
	if title != want {
		t.Errorf("Wrong title found. Want: %s, Got : %s", want, title)
	}
}
