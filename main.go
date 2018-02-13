package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/seankhliao/go3status/protocol"
)

func main() {
	raw, err := ioutil.ReadFile("default.toml")
	var config string
	if err != nil {
		log.Println("file not found")
		config = DefaultConfig
	} else {
		config = string(raw)
	}

	modules, err := ParseConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	s := NewStatus(os.Stdout, protocol.MinimalHeader())
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	if err := s.StartBlocks(modules); err != nil {
		log.Fatal(err)
	}

	for range time.NewTicker(time.Second).C {
		if err := s.Next(); err != nil {
			log.Fatal(err)
		}
	}
}
