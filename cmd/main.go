package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uwezukwechibuzor/QuicksilverCSVQuery/internal/csv"
	delegatorValidatorQuery "github.com/uwezukwechibuzor/QuicksilverCSVQuery/internal/queries/delegatorValidatorMapping"
	validatorDelegatorQuery "github.com/uwezukwechibuzor/QuicksilverCSVQuery/internal/queries/validatorDelegatorMapping"
)

func main() {

	var outputFilename string

	// Root command
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "Quicksilver Data Query App",
	}

	// Command for querying Validator -> Delegator mapping
	queryValidatorDelegatorCmd := &cobra.Command{
		Use:   "query-validator-delegator [validator-address]",
		Short: "Query Validator -> Delegator mapping",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			argValidatorAddress := args[0]

			mapping, err := validatorDelegatorQuery.QueryValidatorDelegatorMapping(argValidatorAddress)

			if err != nil {
				return err
			}
			return csv.WriteCSV(mapping, outputFilename)
		},
	}
	queryValidatorDelegatorCmd.Flags().StringVarP(&outputFilename, "output", "o", "validator_delegator_mapping.csv", "Output filename")

	// Command for querying Delegator -> Validator mapping
	queryDelegatorValidatorCmd := &cobra.Command{
		Use:   "query-delegator-validator [delegator-address]",
		Short: "Query Delegator -> Validator mapping",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			argDelegatorAddress := args[0]

			mapping, err := delegatorValidatorQuery.QueryDelegatorValidatorMapping(argDelegatorAddress)
			if err != nil {
				return err
			}
			return csv.WriteCSV(mapping, outputFilename)
		},
	}

	queryDelegatorValidatorCmd.Flags().StringVarP(&outputFilename, "output", "o", "delegator_validator_mapping.csv", "Output filename")

	// Add commands to root command
	rootCmd.AddCommand(queryValidatorDelegatorCmd)
	rootCmd.AddCommand(queryDelegatorValidatorCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
