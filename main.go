package main

import (
	"log"
	"os"
	"time"

	"github.com/seankhliao/go3status/protocol"
	"github.com/seankhliao/go3status/status"
)

func main() {
	modules, err := ParseConfig("default.toml")
	if err != nil {
		log.Fatal(err)
	}

	s := status.NewStatus(os.Stdout, protocol.MinimalHeader())
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	s.StartBlocks(modules)

	for range time.NewTicker(time.Second).C {
		if err := s.Next(); err != nil {
			log.Fatal(err)
		}
	}
}
