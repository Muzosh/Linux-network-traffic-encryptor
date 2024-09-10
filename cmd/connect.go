/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "<short desc>",
	Long:  `<long desc>`,
	// PreRun: func(cmd *cobra.Command, args []string) {

	// },
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

var otherGateway string

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.LocalFlags().StringVarP(&otherGateway, "gateway", "g", "", "Public IP of other gateway")
	connectCmd.MarkFlagRequired("gateway")
}
