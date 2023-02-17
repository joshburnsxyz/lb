package util

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joshburnsxyz/lb/backend"
	"github.com/joshburnsxyz/lb/serverpool"
)

func ReadBackendsFile(filep string, serverPool *serverpool.ServerPool) {
	file, err := os.Open(filep)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		url, err := url.Parse(line)
		if err != nil {
			log.Fatal(err)
		}
		backend := backend.New(url)
		serverPool.AddBackend(backend)
		fmt.Printf("Loaded backend %s\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
