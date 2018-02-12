package mod

import (
	"time"

	"github.com/seankhliao/go3status/protocol"
)

const MaxInt = int(^uint(0) >> 1)

type Module interface {
	Start() (chan time.Time, chan *protocol.Block)
	NewBlock(t time.Time) *protocol.Block
}

type Mod struct {
	// Internal
	name    string
	counter int
	tick    chan time.Time
	tock    chan *protocol.Block

	// General
	Instance  string
	Frequency int
}

func NewMod(name string, freq int) Mod {
	return Mod{
		name: name,
		tick: make(chan time.Time, 1),
		tock: make(chan *protocol.Block, 1),

		Frequency: freq,
	}
}

func (m *Mod) Start(newBlock func(time.Time) *protocol.Block) (chan time.Time, chan *protocol.Block) {
	go func() {
		for t := range m.tick {
			if m.counter%m.Frequency == 0 {
				m.tock <- newBlock(t)
			}
			m.counter++
		}
	}()
	return m.tick, m.tock
}
