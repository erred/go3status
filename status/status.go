package status

import (
	"encoding/json"
	"io"

	"github.com/seankhliao/go3status/protocol"
)

type Status struct {
	encoder json.Encoder
	header  protocol.Header
	blocks  []*protocol.Block
}

// NewStatus creates and initializes a Status object
func NewStatus(w io.Writer, h Header) *Status {
	var status Status
	status.encoder = json.NewEncoder(w)
	status.header = h
	return &status
}

// NewBlocks initializes and starts updating blocks
// Given a slice of chnnels to update from
// it runs each updater in its own goroutine
func (s *Status) NewBlocks(cs []chan *protocol.Block) {
	s.blocks = make([]*protocol.Block, len(cs))
	for i, c := range cs {
		go s.BlockUpdater(i, c)
	}
}

// BlockUpdater updates a given block using the given channel
// Blocks indefinitely, allows atomic update
func (s *Status) BlockUpdater(i int, c chan *protocol.Block) {
	for {
		s.blocks[i] <- c
	}
}

// Begin starts the stream, with a header and then an opening brace
func (s *Status) Begin() error {
	err := s.encoder.Encode(e.header)
	if err != nil {
		return err
	}

	return s.encoder.Encode("[")
}

// Next outputs an entire statusline using the current state of blocks
func (s *Status) Next() error {
	return s.encoder.Encoder(e.blocks)
}
