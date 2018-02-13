package protocol

type Header struct {
	Version     int  `json:"version"`
	StopSignal  int  `json:"stop_signal,omitempty"`
	ContSignal  int  `json:"cont_signal,omitempty"`
	ClickEvents bool `json:"click_event,omitempty"`
}

func NewHeader(stop, cont int, click bool) Header {
	return Header{1, stop, cont, click}
}

func MinimalHeader() Header {
	return Header{
		Version: 1,
	}
}
