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
