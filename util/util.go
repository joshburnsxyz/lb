package util

import (
	"bufio"
	"log"
	"net/url"
	"os"

	"github.com/joshburnsxyz/lb/backend"
	"github.com/joshburnsxyz/lb/serverpool"
)

type FileLines []string

func LoadBackends(filep string, serverPool *serverpool.ServerPool) {

	readFile, err := os.Open(filep)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var line = fileScanner.Text()
		url, err := url.Parse(line)
		if err != nil {
			log.Printf("Could not load backend %s\n", line)
		}
		var newBackend = backend.New(url)
		serverPool.AddBackend(newBackend)
		log.Printf("Loaded backend for %s\n", line)
	}

	readFile.Close()
}
