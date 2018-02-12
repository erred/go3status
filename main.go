package main

import (
	"log"
	"os"
	"time"

	"github.com/seankhliao/go3status/protocol"
	"github.com/seankhliao/go3status/status"
)

type Module interface {
	Rename(name, instance string)
	Start() (chan time.Time, chan *protocol.Block)
}

func main() {
	conf, err := ParseConfig("default.toml")
	if err != nil {
		log.Fatal(err)
	}

	s := status.NewStatus(os.Stdout, protocol.MinimalHeader())
	if err := s.Begin(); err != nil {
		log.Fatal(err)
	}

	ts, bs := conf.StartBlocks()
	s.RegisterBlocks(bs)

	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		for _, s := range ts {
			s <- t
		}
		if err := s.Next(); err != nil {
			log.Fatal(err)
		}
	}
}
