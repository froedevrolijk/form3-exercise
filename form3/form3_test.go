package form3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if baseUrl := c.BaseUrl.String(); baseUrl != getBaseUrl() {
		t.Errorf("NewClient BaseUrl; got %v, want %v", baseUrl, defaultBaseUrl)
	}

	if client := c.Client(); client != http.DefaultClient {
		t.Errorf("NewClient Client; got %+v, want %+v", client, http.DefaultClient)
	}

	httpTestClient := new(http.Client)
	c = NewClient(httpTestClient)
	if client := c.Client(); client != httpTestClient {
		t.Errorf("NewClient Test Client; got %+v, want %+v", client, httpTestClient)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	type Health struct {
		Status string `json:"status"`
	}

	wantUrl := getBaseUrl() + "/v1/health"
	wantBodyStr := "{\"status\":\"up\"}\n"

	req, _ := c.NewRequest(ctx, http.MethodGet, "/v1/health", &Health{Status: "up"})
	if u := req.URL.String(); u != wantUrl {
		t.Errorf("NewRequest URL; got %v, want %v", u, wantUrl)
	}

	body, _ := io.ReadAll(req.Body)

	if body := string(body); body != wantBodyStr {
		t.Errorf("NewRequest Body; got %v, want %v", body, wantBodyStr)
	}
}

func TestNewRequest_BadUrl(t *testing.T) {
	c := NewClient(nil)

	_, err := c.NewRequest(ctx, http.MethodGet, "/!%$@!", nil)

	if err == nil {
		t.Errorf("Expected an error")
	}

	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected a URL parsing error, got %v", err)
	}
}

func TestNewRequest_EmptyBody(t *testing.T) {
	c := NewClient(nil)

	req, err := c.NewRequest(ctx, http.MethodGet, "/", nil)

	if err != nil {
		t.Errorf("NewRequest returned an error %v", err)
	}

	if req.Body != nil {
		t.Fatalf("Request contains a non-nil Body")
	}
}

func TestNewRequest_InvalidJSON(t *testing.T) {
	c := NewClient(nil)

	type invalidJson struct {
		Invalid map[interface{}]interface{}
	}

	_, err := c.NewRequest(ctx, http.MethodGet, "/", new(invalidJson))

	if err == nil {
		t.Errorf("NewRequest expected an error")
	}

	if err, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Errorf("Expected a JSON error, got %v", err)
	}
}

func TestSendRequest(t *testing.T) {
	teardown := setup()
	defer teardown()

	type TestType struct {
		Field string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		equal(t, r.Method, http.MethodGet)
		fmt.Fprint(w, `{"Field":"v"}`)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)

	body := new(TestType)

	_, err := client.SendRequest(req, body)
	if err != nil {
		t.Fatalf("SendRequest returned an error: %v", err)
	}

	want := &TestType{"v"}

	if !cmp.Equal(want, body) {
		t.Error(cmp.Diff(want, body))
	}
}

func TestSendRequest_HttpError(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad request", http.StatusBadRequest)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)
	_, err := client.SendRequest(req, nil)

	if err == nil {
		t.Error("Expected an HTTP error")
	}
}

func TestSendRequest_NoContent(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	body := new(json.RawMessage)

	req, _ := client.NewRequest(ctx, http.MethodDelete, "/", nil)
	_, err := client.SendRequest(req, body)

	if err != nil {
		t.Fatalf("SendRequest returned an error: %v", err)
	}
}
