package queries

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/uwezukwechibuzor/QuicksilverCSVQuery/internal/api"
)

// Function to query Validator -> Delegator mapping
func QueryInterchainStaking(chainID string) ([]Receipts, error) {
	// Construct the URL with the validator address
	url := fmt.Sprintf(api.InterchainStakingURL, chainID)

	// Make HTTP request to the endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query InterchainStaking Data: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to query InterchainStaking Data: %s", resp.Status)
	}

	// Parse the response body and extract the queried data
	var receiptsResponse ReceiptsResponse

	err = json.NewDecoder(resp.Body).Decode(&receiptsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	// Create a new slice to store the pending receipts
	newReceipts := make([]Receipts, 0)

	// Query all pending receipts in the x/interchainstaking module
	foundNullReceipts := false
	for _, receipt := range receiptsResponse.Receipts {
		if receipt.Completed == "" {
			newReceipts = append(newReceipts, receipt)
			foundNullReceipts = true
		}
	}

	// Check if any receipts were found
	if !foundNullReceipts {
		return nil, fmt.Errorf("no matching accounts found")
	}

	return newReceipts, nil
}
