package protocol

type Header struct {
	Version     int  `json:"version"`
	StopSignal  int  `json:"stop_signal,omitempty"`
	ContSignal  int  `json:"cont_signa,omitemptyl"`
	ClickEvents bool `json:"click_event,omitemptys"`
}

// func NewHeader(stop, cont int, click bool) Header {
// 	return Header{1, stop, cont, click}
// }

// func (h *Header) MarshalJSON([]byte, string) {
// 	return json.Marshal(h)
// }
