package protocol

// Header encodes protocol version and configuration
type Header struct {
	Version     int  `json:"version"`
	StopSignal  int  `json:"stop_signal,omitempty"`
	ContSignal  int  `json:"cont_signal,omitempty"`
	ClickEvents bool `json:"click_event,omitempty"`
}

// NewHeader creates a header with given config
func NewHeader(stop, cont int, click bool) Header {
	return Header{1, stop, cont, click}
}

// MinimalHeader creates a header with only te protocol version
func MinimalHeader() Header {
	return Header{
		Version: 1,
	}
}
