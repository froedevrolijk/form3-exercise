# Form3 Take Home Exercise

[![build](https://github.com/froedevrolijk/form3-exercise/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/froedevrolijk/form3-exercise/actions/workflows/build.yaml) [![codecov](https://codecov.io/gh/froedevrolijk/form3-exercise/branch/dev/graph/badge.svg?token=QDHQMIWDRO)](https://codecov.io/gh/froedevrolijk/form3-exercise)

Go client for accessing the [Form3 account API](https://www.api-docs.form3.tech/api/schemes/bacs/accounts/overview)

### Requirements
* Go 1.18+
* Docker Desktop

### Project setup
#### Build project and run tests:
`docker-compose up`

#### Run linter:
`make lint`

#### Run formatter:
`make fmt`

#### Run tests and show code coverage as HTML:
`make test`  
`go tool cover -html=coverage.out`

#### Run tests using Docker and subsequently stop containers:
`make docker-test`

### Use the client library
#### Create a client:
```go
import (
   "context"
   "fmt"

   "github.com/froedevrolijk/form3-exercise/form3"
)

c := form3.NewClient()
```

#### Create an account:
```go
someUUID := "ad27e265-9605-4b4b-a0e5-3003ea9cc4de"

someAccount := &form3.Account{
	Data: &form3.AccountData{
		Type:           "accounts",
		ID:             someUUID,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Attributes: &form3.AccountAttributes{
			Country:                 "GB",
			BaseCurrency:            "GBP",
			BankID:                  "400300",
			BankIDCode:              "GBDSC",
			Bic:                     "NWBKGB22",
			Name:                    []string{"Samantha Holder"},
			AlternativeNames:        []string{"Sam Holder"},
			AccountClassification:   "Personal",
			JointAccount:            false,
			AccountMatchingOptOut:   false,
			SecondaryIdentification: "A1B2C3D4",
		},
	},
}

account, _, _ := c.Accounts.CreateAccount(context.Background(), someAccount)

fmt.Printf("Account: %+v", account.Account)
```

#### Fetch an account:
```go
account, _, _ := c.Accounts.GetAccount(context.Background(), someUUID)

fmt.Printf("Account: %+v", account.Account)
```

#### Delete an account:
```go
delOpt := &form3.DeleteOptions{
	AccountID: someUUID,
	Version:   0,
}

c.Accounts.DeleteAccount(context.Background(), delOpt)
```

### Note
Author: Froede Vrolijk  
I started learning Go about half a year ago. I'm using Scala at my current job.  