package queries

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/uwezukwechibuzor/QuicksilverCSVQuery/internal/api"
)

// Function to query Delegator -> Validator mapping
func QueryDelegatorValidatorMapping(delegatorAddr string) ([]Validator, error) {
	// Construct the URL with the delegator address
	url := fmt.Sprintf(api.DelegatorValidatorURL, delegatorAddr)

	// Make HTTP request to the endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query Delegator -> Validator mapping: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to query Delegator -> Validator mapping: %s", resp.Status)
	}

	// Parse the response body and extract the queried data
	var mappingResponse DelegatorValidatorMappingResponse
	err = json.NewDecoder(resp.Body).Decode(&mappingResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	return mappingResponse.Mappings, nil
}
