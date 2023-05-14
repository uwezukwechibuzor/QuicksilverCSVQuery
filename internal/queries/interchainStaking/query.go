package queries

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/uwezukwechibuzor/QuicksilverCSVQuery/internal/api"
)

// Function to query Validator -> Delegator mapping
func QueryInterchainStaking(chain_id string) ([]Receipts, error) {
	// Construct the URL with the validator address
	url := fmt.Sprintf(api.InterchainStakingURL, chain_id)

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
	var Receipt ReceiptsResponse

	err = json.NewDecoder(resp.Body).Decode(&Receipt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	return Receipt.Receipts, nil
}
