package form3

import (
	"time"
)

// Account represents an account.
type Account struct {
	Data *AccountData `json:"data,omitempty"`
}

// AccountData represents data related to an account.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        int64              `json:"version,omitempty"`
}

// AccountAttributes represents account attributes for an account.
type AccountAttributes struct {
	AccountClassification   string   `json:"account_classification,omitempty"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 string   `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            bool     `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  string   `json:"status,omitempty"`
	Switched                bool     `json:"switched,omitempty"`
}

// AccountResponse represents a newly created account.
type AccountResponse struct {
	Account *AccountResponseData `json:"data"`
	Self    string               `json:"self"`
}

// AccountResponseData represents data related to an newly created account.
type AccountResponseData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	CreatedOn      time.Time          `json:"created_on"`
	ID             string             `json:"id"`
	ModifiedOn     time.Time          `json:"modified_on"`
	OrganisationID string             `json:"organisation_id"`
	Type           string             `json:"type"`
	Version        int                `json:"version"`
}

// DeleteOptions represents URL parameters for the DELETE endpoint.
type DeleteOptions struct {
	AccountID string
	Version   int64
}
