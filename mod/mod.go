package mod

import (
	"time"

	"github.com/seankhliao/go3status/protocol"
)

const MaxInt = int(^uint(0) >> 1)

type Module interface {
	Start(blocks []*protocol.Block, pos int) error
	NewBlock(t time.Time) *protocol.Block
}

type Mod struct {
	// Internal
	name    string
	counter int
	tick    chan time.Time
	tock    chan *protocol.Block
	blocks  []*protocol.Block
	pos     int

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

func (m *Mod) Start(blocks []*protocol.Block, pos int, blockFunc func(time.Time) *protocol.Block) {
	m.blocks = blocks
	m.pos = pos
	go func() {
		m.blocks[m.pos] = blockFunc(time.Now())
		if m.Frequency > 0 {
			for t := range time.NewTicker(time.Second * time.Duration(m.Frequency)).C {
				m.blocks[m.pos] = blockFunc(t)
			}
		}
	}()
}
