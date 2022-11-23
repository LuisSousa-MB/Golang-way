package main

import "github.com/spf13/cobra"


func main() {
	rootCmd := &amp
	cobra.Command{}
	rootCmd.AddCommand(Hello())

	rootCmd.Execute()
}
