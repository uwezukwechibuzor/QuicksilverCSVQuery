package csv

import (
	"encoding/csv"
	"os"

	"github.com/gocarina/gocsv"
)

// Function to write data to CSV file
func WriteCSV(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := gocsv.MarshalCSV(data, writer); err != nil {
		return err
	}

	return nil
}
