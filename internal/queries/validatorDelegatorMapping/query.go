package queries

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/uwezukwechibuzor/QuicksilverCSVQuery/internal/api"
)

// Function to query Validator -> Delegator mapping
func QueryValidatorDelegatorMapping(validatorAddr string) ([]DelegationResponses, error) {
	// Construct the URL with the validator address
	url := fmt.Sprintf(api.ValidatorDelegatorURL, validatorAddr)

	// Make HTTP request to the endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query Validator -> Delegator mapping: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to query Validator -> Delegator mapping: %s", resp.Status)
	}

	// Parse the response body and extract the queried data
	var mappingResponse ValidatorDelegationResponse

	err = json.NewDecoder(resp.Body).Decode(&mappingResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	return mappingResponse.Mappings, nil
}
