/*
Copyright © 2023 Josh Burns
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joshburnsxyz/lb/backend"
	"github.com/joshburnsxyz/lb/serverpool"
	"github.com/spf13/cobra"
)

var (
	tlsMode     bool
	port        int
	tlsCertPath string
	tlsKeyPath  string
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
		bu1, _ := url.Parse("http://mybackend.com/1")
		b1 := backend.Backend{URL: bu1}
		serverPool.AddBackend(&b1)

		// Fire-off healthcheck sub-routine
		go healthCheck(serverPool)

		// Assign handler and boot server
		portFormat := fmt.Sprintf(":%d", port)
		http.HandleFunc("/", serverPool.Proxy)
		if tlsMode {
			http.ListenAndServeTLS(portFormat, tlsCertPath, tlsKeyPath, nil)
		} else {
			http.ListenAndServe(portFormat, nil)
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
	rootCmd.Flags().IntVarP(&port, "port", "p", 80, "Port to bind too.")
}
