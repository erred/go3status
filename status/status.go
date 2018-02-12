package status

import (
	"encoding/json"
	"io"

	"github.com/seankhliao/go3status/protocol"
)

type Status struct {
	encoder *json.Encoder
	w       io.Writer
	header  protocol.Header
	blocks  []*protocol.Block
}

// NewStatus creates and initializes a Status object
func NewStatus(w io.Writer, h protocol.Header) *Status {
	var status Status
	status.encoder = json.NewEncoder(w)
	status.w = w
	status.header = h
	return &status
}

// RegisterBlocks connects feed channels to their output state
// Given a slice of chnnels to update from
// it runs each updater in its own goroutine
func (s *Status) RegisterBlocks(cs []chan *protocol.Block) {
	s.blocks = make([]*protocol.Block, len(cs))
	for i, c := range cs {
		// go s.BlockUpdater(i, c)
		go func(i int, c chan *protocol.Block) {
			s.blocks[i] = <-c
		}(i, c)
	}
}

// BlockUpdater updates a given block using the given channel
// Blocks indefinitely, allows atomic update
// func (s *Status) BlockUpdater(i int, c chan *protocol.Block) {
// 	for {
// 		s.blocks[i] = <-c
// 	}
// }

// Begin starts the stream, with a header and then an opening brace
func (s *Status) Begin() error {
	if err := s.encoder.Encode(s.header); err != nil {
		return err
	}

	// because [ is not valid json
	_, err := s.w.Write([]byte("["))
	return err
}

// Next outputs an entire statusline using the current state of blocks
func (s *Status) Next() error {
	return s.encoder.Encode(s.blocks)
}
