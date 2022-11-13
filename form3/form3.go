package form3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	defaultBaseUrl = "BASE_URL"
	defaultTimeout = 10 * time.Second
)

// Client manages communication with the Form3 API.
type Client struct {
	BaseUrl  *url.URL
	client   *http.Client
	Accounts *AccountsService
}

type service struct {
	client *Client
}

// NewClient returns a new Form3 API client.
func NewClient() *Client {
	baseUrl, _ := url.Parse(getBaseUrl())

	c := &Client{
		BaseUrl: baseUrl,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}

	c.Accounts = &AccountsService{client: c}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseUrl.ResolveReference(parsedUrl).String()

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u, buf)
	if err != nil {
		return nil, err
	}

	addHeaders(req)

	return req, nil
}

// SendRequest sends an API request and returns the API response.
func (c *Client) SendRequest(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	data, err := CheckResponse(resp)
	if err != nil {
		return response, err
	}

	responseDataReader := bytes.NewReader(data)
	err = json.NewDecoder(responseDataReader).Decode(v)
	if err == io.EOF {
		err = nil
	}

	return response, err
}

// CheckResponse checks the API response for errors.
// Status codes outside the 200 range are considered errors.
func CheckResponse(r *http.Response) ([]byte, error) {
	var errResp = new(ErrorResponse)

	data, err := io.ReadAll(r.Body)

	if code := r.StatusCode; 200 <= code && code <= 299 || code == 308 {
		return data, nil
	}

	if err == nil && data != nil && json.Valid(data) {
		err = json.Unmarshal(data, errResp)
		if err != nil {
			return data, err
		}
	}

	errResp.Status = r.StatusCode

	return data, errResp
}

// ErrorResponse represents an error response with a status code
// and an error message.
type ErrorResponse struct {
	Status       int
	ErrorMessage string `json:"error_message,omitempty"`
}

// Error formats the ErrorResponse.
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Status: %d Message: %v", e.Status, e.ErrorMessage)
}

func addHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Host", getBaseUrl())
	req.Header.Set("Date", time.Now().Format(time.RFC1123))
}

func getBaseUrl() string {
	u := os.Getenv(defaultBaseUrl)
	if u == "" {
		u = "http://localhost:8080"
	}

	return u
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	resp := &Response{Response: r}
	return resp
}
