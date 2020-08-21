package jsonutils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/BenSchoeggl/mongo-challenge/testutils"
	"github.com/stretchr/testify/require"
)

func TestFlatten(t *testing.T) {
	testCases := testutils.GetTestFileData(t)
	testResults := map[string][]byte{}
	for caseName, caseData := range testCases {
		data := map[string]interface{}{}
		require.Nil(t, json.Unmarshal(caseData, &data), "could not unmarshal test data to map")
		fmt.Println("testing data: ", data)
		flattenedJSON := Flatten(data)
		flattenedJSONBytes, err := json.MarshalIndent(flattenedJSON, "", "    ")
		require.Nil(t, err, "could marshal test data to bytes")
		testResults[caseName] = flattenedJSONBytes
	}
	testutils.ValidateResults(t, testResults)
}
