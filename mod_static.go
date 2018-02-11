package main

import (
	"time"

	"github.com/seankhliao/go3status/protocol"
)

type ModStatic struct {
	name       string
	instance   string
	Text       string `toml:"text"`
	Color      string
	Background string
	Border     string
}

func (m *ModStatic) Start() (chan time.Time, chan *protocol.Block) {
	var block protocol.Block

	block.FullText = m.Text
	block.Color = m.Color
	block.Background = m.Background
	block.Border = m.Border

	t := make(chan time.Time)
	go func() {
		for {
			<-t

		}
	}()

	c := make(chan *protocol.Block, 1)
	c <- &block
	return t, c
}

func (m *ModStatic) Rename(name, instance string) {
	m.name = name
	m.instance = instance
}
