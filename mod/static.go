package mod

import (
	"time"

	"github.com/seankhliao/go3status/protocol"
)

type Static struct {
	Mod

	Text  string
	Color string
}

func NewStatic() Module {
	return &Static{
		Mod: NewMod("static", 0),
	}
}

func (m *Static) NewBlock(t time.Time) *protocol.Block {
	return &protocol.Block{
		FullText: m.Text,
		Color:    m.Color,
		Name:     m.name,
		Instance: m.Instance,
	}
}

func (m *Static) Start(blocks []*protocol.Block, pos int) error {
	m.Mod.Start(blocks, pos, m.NewBlock)
	return nil
}
