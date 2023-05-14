package queries

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/uwezukwechibuzor/QuicksilverCSVQuery/internal/api"
)

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

	return mappingResponse.Accounts, nil
}
