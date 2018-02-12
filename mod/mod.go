package mod

import (
	"time"

	"github.com/seankhliao/go3status/protocol"
)

type Module interface {
	Start() (chan time.Time, chan *protocol.Block)
}
