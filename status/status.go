package status

import (
	"encoding/json"
	"io"

	"github.com/seankhliao/go3status/mod"
	"github.com/seankhliao/go3status/protocol"
)

type Status struct {
	encoder *json.Encoder
	w       io.Writer
	header  protocol.Header
	Blocks  []*protocol.Block
}

// NewStatus creates and initializes a Status object
func NewStatus(w io.Writer, h protocol.Header) *Status {
	var status Status
	status.encoder = json.NewEncoder(w)
	status.w = w
	status.header = h
	return &status
}

func (s *Status) StartBlocks(modules []mod.Module) {
	s.Blocks = make([]*protocol.Block, len(modules))
	for i, module := range modules {
		module.Start(s.Blocks, i)
	}
}

// Start starts the stream, with a header and then an opening brace
func (s *Status) Start() error {
	if err := s.encoder.Encode(s.header); err != nil {
		return err
	}

	// because [ is not valid json
	_, err := s.w.Write([]byte("["))
	return err
}

// Next outputs an entire statusline using the current state of blocks
func (s *Status) Next() error {
	return s.encoder.Encode(s.Blocks)
}
