package mod

import (
	"time"

	"github.com/seankhliao/go3status/protocol"
)

type Static struct {
	name     string
	instance string

	Text       string `toml:"text"`
	Color      string
	Background string
	Border     string

	tick chan time.Time
	tock chan *protocol.Block
	// <= 0: never update
	// > 0: update every n ticks (seconds)
	Frequency int
	counter   int
}

func NewStatic(instance string) Module {
	var m Static

	m.instance = instance
	m.tick = make(chan time.Time, 1)
	m.tock = make(chan *protocol.Block, 1)

	// defaults
	m.name = "static"
	m.Frequency = 0
	return &m
}

func (m *Static) NewBlock(t time.Time) *protocol.Block {
	var block protocol.Block

	block.FullText = m.Text
	block.Color = m.Color
	block.Background = m.Background
	block.Border = m.Border

	return &block
}

func (m *Static) Start() (chan time.Time, chan *protocol.Block) {
	go func() {
		if m.Frequency <= 0 {
			m.tock <- m.NewBlock(time.Now())
		}
		for t := range m.tick {
			m.counter++
			if m.counter == m.Frequency {
				m.counter = 0
				m.tock <- m.NewBlock(t)
			}
		}
	}()
	return m.tick, m.tock
}
