package form3

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const testdataPath = "../testdata/"

func readFixture(path string) string {
	b, err := os.ReadFile(testdataPath + path)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func equal[T comparable](t *testing.T, want, got T) {
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func equalRequestBody[T comparable](t *testing.T, r *http.Request, want, got T) {
	err := json.NewDecoder(r.Body).Decode(got)
	if err != nil {
		panic(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

type apiResponse interface {
	*Account | *AccountResponse
}

func createApiResponse[T apiResponse](fileName string) T {
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var data T

	err = json.Unmarshal([]byte(b), &data)
	if err != nil {
		panic(err)
	}

	return data
}
