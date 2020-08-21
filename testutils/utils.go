package testutils

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

// GetTestFileData returns an array of byte arrays of the test files for the calling
// unit test function. The input test files should be stored at the following path:
// <directory of calling unit test file>/test-data/input/<name of calling unit test>
func GetTestFileData(t *testing.T) map[string][]byte {
	testFileDir := path.Join(getTestFileDir(t), "input", t.Name())
	files, err := ioutil.ReadDir(testFileDir)
	if err != nil {
		t.Errorf("error when trying to read input test directory: %v", err)
	}
	result := map[string][]byte{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileBytes, err := ioutil.ReadFile(path.Join(testFileDir, file.Name()))
		if err != nil {
			t.Errorf("error when trying to read test file: %v", err)
		}
		result[file.Name()] = fileBytes
	}
	return result
}

// ValidateResults compares the results of a unit test to the golden data on disk. If the input data
// differs from the data on disk, the currently running test will be failed now, and the data on
// disk will be overwritten with the incoming data.
// The golden data should be located at
// <path to calling file>/test_data/golden_data/<test name>/<test case names>
func ValidateResults(t *testing.T, results map[string][]byte) {
	goldenDataDir := path.Join(getTestFileDir(t), "golden_data", t.Name())
	require.Nil(t, os.MkdirAll(goldenDataDir, 0700))
	files, err := ioutil.ReadDir(goldenDataDir)
	require.Nil(t, err, "could not read golden data dir")
	existingGoldenFiles := map[string][]byte{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := path.Join(goldenDataDir, file.Name())
		fileBytes, err := ioutil.ReadFile(filePath)
		require.Nil(t, err, "could not read golden data file at path: %v", filePath)
		existingGoldenFiles[filePath] = fileBytes
	}
	var goldenDataChanged bool
	for caseName, caseResultData := range results {
		caseFilePath := path.Join(goldenDataDir, caseName)
		existingGoldenData, ok := existingGoldenFiles[caseFilePath]
		// if it doesn't already exist, or has changed, we need to write
		if !ok || !bytes.Equal(existingGoldenData, caseResultData) {
			goldenDataChanged = true
			require.Nil(
				t,
				ioutil.WriteFile(caseFilePath, caseResultData, 0644),
				"couldn't write golden data",
			)
		}
		// deleting from a map while youre iterating is safe in go
		delete(existingGoldenFiles, caseFilePath)
	}
	// if there is anything left in existingGoldenFiles, we need to delete it
	for noLongerNeededGoldenFile := range existingGoldenFiles {
		goldenDataChanged = true
		require.Nil(t, os.Remove(noLongerNeededGoldenFile), "could not remove old golden data")
	}
	require.False(t, goldenDataChanged, "golden data changed, check diff")
}

func getTestFileDir(t *testing.T) string {
	_, file, _, ok := runtime.Caller(2)
	require.True(t, ok, "could not get caller information to retreive input data")
	return path.Join(path.Dir(file), "test_data")
}
