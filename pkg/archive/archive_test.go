package archive

import (
	"testing"
)

func TestGetArchives(t *testing.T) {
	testFiles := []string{"/foo/bar/2020-01-02T13:00:01", "/foo/bar/2019-01-02T13:00:01"}
	archives := buildArchives("./StashDB", testFiles)
	expected := archive{"/foo/bar", []string{"2020-01-02T13:00:01", "2019-01-02T13:00:01"}}
	if archives[0].Url != expected.Url {
		t.Errorf("expected: %s actual: %s", expected.Url, archives[0].Url)
	}
	if archives[0].Dates[0] != expected.Dates[0] && len(archives[0].Dates) != len(expected.Dates) {
		t.Errorf("expected: %s actual: %s", expected.Dates, archives[0].Dates)
	}
}
