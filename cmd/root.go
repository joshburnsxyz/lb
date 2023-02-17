/*
Copyright © 2023 Josh Burns
*/
package cmd

import (
	"os"

	"github.com/joshburnsxyz/lb/serverpool"
	"github.com/joshburnsxyz/lb/util"
	"github.com/joshburnsxyz/lb/server"
	"github.com/spf13/cobra"
)

var (
	backendsList string
)

var rootCmd = &cobra.Command{
	Use:   "lb",
	Short: "Expiremental HTTP Load Balancer",
	Run: func(cmd *cobra.Command, args []string) {
		serverPool := serverpool.New()
		server := server.New(serverPool, 8188)

		// Load backends into server pool
		util.LoadBackends(backendsList, serverPool)
		
		// Launch server
		server.ListenAndServe()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&backendsList, "backends", "b", "/etc/lb/backends.txt", "List of backends")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
