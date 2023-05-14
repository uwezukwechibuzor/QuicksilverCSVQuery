package queries

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/uwezukwechibuzor/QuicksilverCSVQuery/internal/api"
)

// Function to query Vesting Accounts
// Function to query Vesting Accounts
func QueryVestingAccounts() ([]PeriodicVestingAccount, error) {
	// Construct the URL with the account query
	url := fmt.Sprintf(api.AccountsURL)

	// Make HTTP request to the endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query accounts: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to query accounts: %s", resp.Status)
	}

	// Parse the response body and extract the queried data
	var mappingResponse AccountsResponse

	err = json.NewDecoder(resp.Body).Decode(&mappingResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	// Create a new slice to store the filtered accounts
	filteredAccounts := make([]PeriodicVestingAccount, 0)

	// Check that the account type returned are DelayedVestingAccount, PeriodicVestingAccount, and PermanentLockedAccount
	for _, account := range mappingResponse.Accounts {
		if account.Type == "/cosmos.vesting.v1beta1.PeriodicVestingAccount" || account.Type == "/cosmos.vesting.v1beta1.DelayedVestingAccount" || account.Type == "/cosmos.vesting.v1beta1.PermanentLockedAccount" {
			filteredAccounts = append(filteredAccounts, account)
		}
	}
	log.Println(filteredAccounts)

	return filteredAccounts, nil
}
