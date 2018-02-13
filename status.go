package main

import (
	"io"

	"github.com/seankhliao/go-i3bar-protocol"
	"github.com/seankhliao/go3status/mod"
)

type Status struct {
	encoder *protocol.Encoder
	header  protocol.Header
	Blocks  []*protocol.Block
}

// NewStatus creates and initializes a Status object
func NewStatus(w io.Writer, h protocol.Header) *Status {
	var status Status
	status.encoder = protocol.NewEncoder(w)
	status.header = h
	return &status
}

func (s *Status) StartBlocks(modules []mod.Module) error {
	s.Blocks = make([]*protocol.Block, len(modules))
	for i, module := range modules {
		if err := module.Start(s.Blocks, i); err != nil {
			return err
		}
	}
	return nil
}

// Start starts the stream, with a header and then an opening brace
func (s *Status) Start() error {
	if err := s.encoder.Encode(s.header); err != nil {
		return err
	}
	return s.encoder.BeginArray()
}

// Next outputs an entire statusline using the current state of blocks
func (s *Status) Next() error {
	return s.encoder.Encode(s.Blocks)
}
