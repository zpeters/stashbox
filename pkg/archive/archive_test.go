package archive

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetArchives(t *testing.T) {
	var emptyArchive []Archive

	tests := []struct {
		inputPath        string
		expectedArchives []Archive
		expectedError    error
	}{
		{".", emptyArchive, fmt.Errorf("no archives found in %s", ".")},
	}

	for _, tc := range tests {
		got, err := GetArchives(tc.inputPath)
		if tc.expectedError == nil {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tc.expectedError.Error())
		}
		require.Equal(t, tc.expectedArchives, got)
	}
}

func TestBuildArchives(t *testing.T) {
	tests := []struct {
		inputTestFiles  []string
		expectedArchive []Archive
		expectedError   error
	}{
		{[]string{"/foo/bar/2020-01-02T13:00:01", "/foo/bar/2019-01-02T13:00:01"}, []Archive{{"/foo/bar", []string{"2020-01-02T13:00:01", "2019-01-02T13:00:01"}}}, nil},
		{[]string{"/foo/bar/20fdsafds20-01-02T13:00:01", "/fofdjsakl;o/b2019-01-02T13:00:01"}, []Archive{}, errors.New("Error building archive")},
	}

	for _, tc := range tests {
		got, err := buildArchives("./StashDB", tc.inputTestFiles)
		require.Equal(t, tc.expectedArchive, got)

		if tc.expectedError == nil {
			require.NoError(t, err)
		} else {
			require.Equal(t, tc.expectedError.Error(), err.Error())
		}
	}
}
