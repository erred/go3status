package protocol

type Header struct {
	Version     int  `json:"version"`
	StopSignal  int  `json:"stop_signal,omitempty"`
	ContSignal  int  `json:"cont_signal,omitemptyl"`
	ClickEvents bool `json:"click_event,omitemptys"`
}

func NewHeader(stop, cont int, click bool) Header {
	return Header{1, stop, cont, click}
}

func MinimalHeader() Header {
	var header Header
	header.Version = 1
	return header
}
