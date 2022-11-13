package form3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountsService_Create(t *testing.T) {
	c := NewClient()

	fixture := readFixture("create-account.json")

	accountData := new(Account)

	err := json.Unmarshal([]byte(fixture), accountData)
	if err != nil {
		panic(err)
	}

	someUUID := uuid.NewString()
	accountData.Data.ID = someUUID

	account, resp, err := c.Accounts.CreateAccount(context.Background(), accountData)
	assert.Nil(t, err, "expecting nil err")
	assert.NotNil(t, resp, "expecting non-nil result")
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	want := createApiResponse[*AccountResponse](testdataPath + "account-response.json")
	if !cmp.Equal(want, account, cmpopts.IgnoreFields(AccountResponse{}, "Account.CreatedOn", "Account.ModifiedOn", "Account.ID")) {
		t.Error(cmp.Diff(want, account))
	}
}

func TestAccountsService_Create_EmptyAttribute(t *testing.T) {
	c := NewClient()

	fixture := readFixture("create-account.json")

	accountData := new(Account)

	err := json.Unmarshal([]byte(fixture), accountData)
	if err != nil {
		panic(err)
	}

	accountData.Data.ID = ""

	account, resp, err := c.Accounts.CreateAccount(context.Background(), accountData)
	assert.Nil(t, account, "expecting nil account")
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, err.Error(), "id in body is required")
}

func TestAccountsService_Create_InvalidAttribute(t *testing.T) {
	c := NewClient()

	fixture := readFixture("create-account.json")

	accountData := new(Account)

	err := json.Unmarshal([]byte(fixture), accountData)
	if err != nil {
		panic(err)
	}

	accountData.Data.Attributes.BankID = "invalid"

	account, resp, err := c.Accounts.CreateAccount(context.Background(), accountData)
	assert.Nil(t, account, "expecting nil account")
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, err.Error(), "bank_id in body should match '^[A-Z0-9]{0,16}$'")
}

func TestAccountsService_Get(t *testing.T) {
	c := NewClient()

	fixture := readFixture("create-account.json")

	accountData := new(Account)

	err := json.Unmarshal([]byte(fixture), accountData)
	if err != nil {
		panic(err)
	}

	someUUID := uuid.NewString()
	accountData.Data.ID = someUUID

	_, _, _ = c.Accounts.CreateAccount(context.Background(), accountData)

	account, resp, err := c.Accounts.GetAccount(context.Background(), someUUID)
	assert.Nil(t, err, "expecting nil err")
	assert.NotNil(t, resp, "expecting non-nil result")
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	want := createApiResponse[*AccountResponse](testdataPath + "account-response.json")
	if !cmp.Equal(want, account, cmpopts.IgnoreFields(AccountResponse{}, "Account.CreatedOn", "Account.ModifiedOn", "Account.ID")) {
		t.Error(cmp.Diff(want, account))
	}
}

func TestAccountsService_GetNonExistentId(t *testing.T) {
	c := NewClient()

	someUUID := uuid.NewString()

	account, resp, err := c.Accounts.GetAccount(context.Background(), someUUID)
	assert.Nil(t, account, "expecting nil account")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Contains(t, err.Error(), fmt.Sprintf("record %v does not exist", someUUID))
}

func TestAccountsService_Get_InvalidUUID(t *testing.T) {
	c := NewClient()

	account, resp, err := c.Accounts.GetAccount(context.Background(), "invalid-uuid")
	assert.Nil(t, account, "expecting nil account")
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, err.Error(), "id is not a valid uuid")
}

func TestAccountsService_Delete(t *testing.T) {
	c := NewClient()

	fixture := readFixture("create-account.json")

	accountData := new(Account)

	err := json.Unmarshal([]byte(fixture), accountData)
	if err != nil {
		panic(err)
	}

	someUUID := uuid.NewString()
	accountData.Data.ID = someUUID

	_, _, _ = c.Accounts.CreateAccount(context.Background(), accountData)

	delOptions := &DeleteOptions{
		AccountID: someUUID,
		Version:   0,
	}

	resp, err := c.Accounts.DeleteAccount(context.Background(), delOptions)

	assert.Nil(t, err, "expecting nil error")
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestAccountsService_Delete_NotFound(t *testing.T) {
	c := NewClient()

	delOptions := &DeleteOptions{
		AccountID: uuid.NewString(),
		Version:   0,
	}

	resp, err := c.Accounts.DeleteAccount(context.Background(), delOptions)

	assert.NotNil(t, err, "expecting non-nil error")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAccountsService_Delete_InvalidUUID(t *testing.T) {
	c := NewClient()

	delOptions := &DeleteOptions{
		AccountID: "invalid-uuid",
		Version:   0,
	}

	resp, err := c.Accounts.DeleteAccount(context.Background(), delOptions)

	assert.NotNil(t, err, "expecting non-nil error")
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAccountsService_Delete_InvalidVersion(t *testing.T) {
	c := NewClient()

	delOptions := &DeleteOptions{
		AccountID: uuid.NewString(),
		Version:   5,
	}

	resp, err := c.Accounts.DeleteAccount(context.Background(), delOptions)

	assert.NotNil(t, err, "expecting non-nil error")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
