package main

import (
	"github.com/spf13/cobra"
	"log"
)

func main() {
	var rootCmd = &cobra.Command{Use: "kvsctl"}
	rootCmd.PersistentFlags().StringP("address", "a", "127.0.0.1:2315", "Server address")
	//rootCmd.AddCommand(NewSetCommand(), NewGetCommand(), NewDeleteCommand(), NewMemberCommand())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
