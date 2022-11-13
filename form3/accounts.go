package form3

import (
	"context"
	"fmt"
	"net/http"
)

// AccountsService handles communication with the Account resource methods of the Form3 API.
//
// Form3 API docs: https://www.api-docs.form3.tech/api/schemes/bacs/accounts/resource-types
type AccountsService service

// GetAccount fetches an account.
//
// Form3 API docs: https://www.api-docs.form3.tech/api/schemes/bacs/accounts/fetch-an-account
func (s *AccountsService) GetAccount(ctx context.Context, id string) (*AccountResponse, *Response, error) {
	u := fmt.Sprintf("/v1/organisation/accounts/%v", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	account := new(AccountResponse)

	resp, err := s.client.SendRequest(req, account)
	if err != nil {
		return nil, resp, err
	}

	return account, resp, err
}

// CreateAccount creates an account.
//
// Form3 API docs: https://www.api-docs.form3.tech/api/schemes/bacs/accounts/create-an-account
func (s *AccountsService) CreateAccount(ctx context.Context, body *Account) (*AccountResponse, *Response, error) {
	u := "/v1/organisation/accounts"

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, body)
	if err != nil {
		return nil, nil, err
	}

	account := new(AccountResponse)

	resp, err := s.client.SendRequest(req, account)
	if err != nil {
		return nil, resp, err
	}

	return account, resp, nil
}

// DeleteAccount deletes an account.
//
// Form3 API docs: https://www.api-docs.form3.tech/api/schemes/bacs/accounts/delete-an-account
func (s *AccountsService) DeleteAccount(ctx context.Context, options *DeleteOptions) (*Response, error) {
	u := fmt.Sprintf("/v1/organisation/accounts/%v?version=%v", options.AccountID, options.Version)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.SendRequest(req, nil)
}
