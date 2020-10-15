package crawler

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	// Setup the test environment
	tempDir := os.TempDir()
	archivePath := path.Join(tempDir, "STASHBOX")
	defer os.RemoveAll(archivePath)

	// Setup our crawler
	c, err := NewCrawler(archivePath)
	require.NoError(t, err)

	// Add some urls
	err = c.AddURL("http://google.com")
	require.NoError(t, err)
	err = c.AddURL("https://thehelpfulhacker.net")
	require.NoError(t, err)

	// Crawl the sites
	err = c.Crawl()
	require.NoError(t, err)

	// Save the sites
	err = c.Save()
	require.NoError(t, err)

	// Get the contents of the archivePath on the file system
	files, err := ioutil.ReadDir(archivePath)
	require.NoError(t, err)

	// there should be two domain folders
	require.Len(t, files, 2)

	// TODO add some more sophisticated testing
}

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
	var tests = []struct {
		inputDir       string
		inputURL       string
		expectedOutput string
		expectedError  error
	}{
		{"./StashDB", "http://www.google.com/a/test.html", "StashDB/www.google.com/a/test.html", nil},
		// See https://golang.org/src/net/url/url_test.go "parseRequestURLTests"
		{"./AnotherDB", " http://foo.com", "", errors.New("parse \" http://foo.com\": first path segment in URL cannot contain colon")},
	}
	for _, tt := range tests {
		actual, err := buildPath(tt.inputDir, tt.inputURL)
		require.Equal(t, tt.expectedOutput, actual)
		if tt.expectedError == nil {
			require.NoError(t, err)
		} else {
			require.Equal(t, tt.expectedError.Error(), err.Error())
		}
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
