package form3

import (
	"time"
)

// Account represents an account.
type Account struct {
	Data *AccountData `json:"data"`
}

// AccountData represents data related to an account.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes"`
	ID             string             `json:"id"`
	OrganisationID string             `json:"organisation_id"`
	Type           string             `json:"type"`
	Version        int64              `json:"version"`
}

// AccountAttributes represents account attributes for an account.
type AccountAttributes struct {
	AccountClassification   string   `json:"account_classification"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
	AccountNumber           string   `json:"account_number"`
	AlternativeNames        []string `json:"alternative_names"`
	BankID                  string   `json:"bank_id"`
	BankIDCode              string   `json:"bank_id_code"`
	BaseCurrency            string   `json:"base_currency"`
	Bic                     string   `json:"bic"`
	Country                 string   `json:"country"`
	Iban                    string   `json:"iban"`
	JointAccount            bool     `json:"joint_account"`
	Name                    []string `json:"name"`
	SecondaryIdentification string   `json:"secondary_identification"`
	Status                  *string  `json:"status,omitempty"`
	Switched                bool     `json:"switched"`
}

// AccountResponse represents the response for a fetched account.
type AccountResponse struct {
	Account *AccountResponseData `json:"data"`
	Self    string               `json:"self"`
}

// AccountResponseData represents the response data related to a fetched account.
type AccountResponseData struct {
	Attributes     *AccountAttributes `json:"attributes"`
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
