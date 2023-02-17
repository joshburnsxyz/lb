/*
Copyright © 2023 Josh Burns
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joshburnsxyz/lb/serverpool"
	"github.com/joshburnsxyz/lb/util"
	"github.com/spf13/cobra"
)

var (
	backendsFilePath string
	tlsMode          bool
	port             int
	tlsCertPath      string
	tlsKeyPath       string
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

		// Load backends into server pool
		util.ReadBackendsFile(backendsFilePath, serverPool)
		fmt.Println("All backends loaded")

		// Assign handler and boot server
		portFormat := fmt.Sprintf(":%d", port)
		http.HandleFunc("/", serverPool.Proxy)
		if tlsMode {
			http.ListenAndServeTLS(portFormat, tlsCertPath, tlsKeyPath, nil)
		} else {
			http.ListenAndServe(portFormat, nil)
		}

		// Fire-off healthcheck sub-routine
		go healthCheck(serverPool)
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

	// TCP port to bind too
	rootCmd.Flags().IntVarP(&port, "port", "p", 80, "Port to bind too.")

	// Configure backends file path
	rootCmd.Flags().StringVarP(&backendsFilePath, "backends", "b", "/etc/lb/backends.txt", "List of URLs to backends")
}
