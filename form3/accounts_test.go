package form3

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

const testUUID = "ad27e265-9605-4b4b-a0e5-3003ea9cc4de"

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
	ctx    = context.Background()
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient(nil)

	url, _ := url.Parse(server.URL)
	client.BaseUrl = url

	return func() {
		server.Close()
	}
}

func TestGetAccount(t *testing.T) {
	teardown := setup()
	defer teardown()

	u := fmt.Sprintf("/v1/organisation/accounts/%v", testUUID)

	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		equal(t, r.Method, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, readFixture("account-response.json"))
	})

	account, _, err := client.Accounts.GetAccount(ctx, testUUID)
	if err != nil {
		t.Errorf("GetAccount returned an error: %v", err)
	}

	want := createApiResponse[*AccountResponse](testdataPath + "account-response.json")

	if !cmp.Equal(want, account) {
		t.Error(cmp.Diff(want, account))
	}
}

func TestGetAccount_InvalidUUID(t *testing.T) {
	teardown := setup()
	defer teardown()

	invalidUUID := "InvalidUUID"
	u := fmt.Sprintf("/v1/organisation/accounts/%v", invalidUUID)

	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		equal(t, r.Method, http.MethodGet)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, readFixture("invalid-uuid.json"))
	})

	_, _, err := client.Accounts.GetAccount(ctx, invalidUUID)

	want := &ErrorResponse{
		Status:       400,
		ErrorMessage: "id is not a valid uuid",
	}

	if !cmp.Equal(want, err) {
		t.Error(cmp.Diff(want, err))
	}

}

func TestGetAccount_NotFound(t *testing.T) {
	teardown := setup()
	defer teardown()

	u := fmt.Sprintf("/v1/organisation/accounts/%v", testUUID)

	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		equal(t, r.Method, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, readFixture("not-found.json"))
	})

	_, _, err := client.Accounts.GetAccount(ctx, testUUID)

	want := &ErrorResponse{
		Status:       404,
		ErrorMessage: fmt.Sprintf("record %v does not exist", testUUID),
	}

	if !cmp.Equal(want, err) {
		t.Error(cmp.Diff(want, err))
	}
}

func TestCreateAccount(t *testing.T) {
	teardown := setup()
	defer teardown()

	body := createApiResponse[*Account](testdataPath + "create-account.json")

	mux.HandleFunc("/v1/organisation/accounts", func(w http.ResponseWriter, r *http.Request) {
		equal(t, r.Method, http.MethodPost)
		equalRequestBody(t, r, body, new(Account))
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, readFixture("account-response.json"))
	})

	account, _, err := client.Accounts.CreateAccount(ctx, body)
	if err != nil {
		t.Errorf("CreateAccount returned an error: %v", err)
	}

	want := createApiResponse[*AccountResponse](testdataPath + "account-response.json")

	if !cmp.Equal(want, account) {
		t.Error(cmp.Diff(want, account))
	}
}

func TestDeleteAccount(t *testing.T) {
	teardown := setup()
	defer teardown()

	u := fmt.Sprintf("/v1/organisation/accounts/%v", testUUID)

	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		equal(t, r.Method, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	delOpt := &DeleteOptions{
		AccountID: testUUID,
		Version:   0,
	}

	resp, err := client.Accounts.DeleteAccount(ctx, delOpt)
	if err != nil {
		t.Errorf("DeleteAccount returned an error: %v", err)
	}

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
