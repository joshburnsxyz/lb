/*
Copyright © 2023 Josh Burns
*/
package cmd

import (
	"log"
	"os"
	"time"

	"github.com/joshburnsxyz/lb/server"
	"github.com/joshburnsxyz/lb/serverpool"
	"github.com/joshburnsxyz/lb/util"
	"github.com/spf13/cobra"
)

var (
	backendsList string
	tlsMode      bool
	tlsCertPath  string
	tlsKeyPath   string
)

func healthCheck(serverPool *serverpool.ServerPool) {
	log.Println("Starting health check...")
	serverPool.HealthCheck()
	log.Println("Health check completed")
}

var rootCmd = &cobra.Command{
	Use:   "lb",
	Short: "Expiremental HTTP Load Balancer",
	Run: func(cmd *cobra.Command, args []string) {
		serverPool := serverpool.New()
		server := server.New(serverPool, 8188)

		// Load backends into server pool
		util.LoadBackends(backendsList, serverPool)

		// Fire-off healthcheck sub-routine
		go healthCheck(serverPool)

		// Launch server
		if tlsMode {
			log.Println("Running with TLS/SSL support")
			log.Fatal(server.ListenAndServe())
		} else {
			log.Println("Running without TLS/SSL support")
			log.Fatal(server.ListenAndServeTLS(tlsCertPath, tlsKeyPath))
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// path to list of backends
	rootCmd.Flags().StringVarP(&backendsList, "backends", "b", "/etc/lb/backends.txt", "List of backends")

	// TLS Settings and configuration
	rootCmd.Flags().BoolVarP(&tlsMode, "tls", "t", false, "Run server in TLS (SSL) mode.")
	rootCmd.Flags().StringVarP(&tlsCertPath, "cert", "c", "", "TLS certificate file.")
	rootCmd.Flags().StringVarP(&tlsKeyPath, "key", "k", "", "TLS key file.")
}
