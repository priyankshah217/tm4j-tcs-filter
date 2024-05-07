package cmd

import "testing"

func TestFilterCmd(t *testing.T) {
	rootCmd.SetArgs([]string{"filter", "-f", "/Users/priyank/Downloads/trust-test-cases.csv", "-l", "backend_open,backend_stable"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
