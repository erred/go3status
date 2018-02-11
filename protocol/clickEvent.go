package protocol

type ClickEvent struct {
	Name     string `json:"name"`
	Instance string `json:"instance"`
	Button   int    `json:"button"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

// func (c clickEvent) UnmarshalJSON(b []byte) error {
// 	return json.Unmarshal(b, c)
// }
