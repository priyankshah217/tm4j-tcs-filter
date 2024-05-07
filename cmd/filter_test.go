package cmd

import "testing"

func TestFilterCmd(t *testing.T) {
	rootCmd.SetArgs([]string{"filter", "-f", "path/to/csv", "-l", "backend_open,backend_stable"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
