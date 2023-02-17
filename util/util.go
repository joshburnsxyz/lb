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

type FileLines []string

func LoadBackends(filep string, serverPool serverpool.ServerPool) {

	readFile, err := os.Open(filep)
	if err != nil {
		log.Println(fmt.Sprintln("ERROR: Could not open backend list %s", err.Error()))
		return nil
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var line = fileScanner.Text()
		url, err := url.Parse(line)
		if err != nil {
			log.Println(fmt.Sprintln("Could not load backend %s", line))
		}
		var newBackend = backend.New(url)
		serverPool.AddBackend(newBackend)
		log.Println(fmt.Sprintln("Loaded backend for %s", line))
	}

	readFile.Close()
}
