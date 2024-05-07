package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var labels []string
var fileName string

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		records, err := readCSV(fileName)
		if err != nil {
			return
		}
		testCasesWithoutServiceLabels := filterTestCasesBasedOnLabels(records, labels)
		for _, testCase := range testCasesWithoutServiceLabels {
			cmd.Println(testCase)
		}
	},
}

func filterTestCasesBasedOnLabels(records [][]string, labels []string) []string {
	var filteredRecords [][]string
	var testWithoutServiceLabels []string
	for _, record := range records {
		for _, label := range labels {
			if strings.Contains(record[8], label) {
				filteredRecords = append(filteredRecords, record)
			}
		}
	}
	for _, record := range filteredRecords {
		if !strings.Contains(record[8], "service_") {
			testWithoutServiceLabels = append(testWithoutServiceLabels, record[0])
		}
	}
	return testWithoutServiceLabels
}

/*
readCSV reads the CSV file
*/
func readCSV(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return records, nil
}

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.Flags().StringSliceVarP(&labels, "labels", "l", []string{}, "Filter test cases by the given criteria")
	filterCmd.Flags().StringVarP(&fileName, "file", "f", "", "The file to filter test cases from")
}
