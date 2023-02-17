/*
Copyright © 2023 Josh Burns
*/
package cmd

import (
	"log"
	"os"

	"github.com/joshburnsxyz/lb/server"
	"github.com/joshburnsxyz/lb/serverpool"
	"github.com/spf13/cobra"
)

var (
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
	// TLS Settings and configuration
	rootCmd.Flags().BoolVarP(&tlsMode, "tls", "t", false, "Run server in TLS (SSL) mode.")
	rootCmd.Flags().StringVarP(&tlsCertPath, "cert", "c", "", "TLS certificate file.")
	rootCmd.Flags().StringVarP(&tlsKeyPath, "key", "k", "", "TLS key file.")
}
