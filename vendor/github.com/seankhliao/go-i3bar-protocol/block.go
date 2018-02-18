package protocol

// Alignment of text within a block
type Alignment string

const (
	// Left aligns text to left
	Left Alignment = "left"
	// Right aligns text to right
	Right Alignment = "right"
	// Center aligns text
	Center Alignment = "center"
)

// Block represents a single cell in the status line
type Block struct {
	FullText   string `json:"full_text"`
	ShortText  string `json:"short_text,omitempty"`
	Color      string `json:"color,omitempty"`
	Background string `json:"background,omitempty"`
	Border     string `json:"border,omitempty"`

	// should also allow pixels
	MinWidth string `json:"min_width,omitempty"`
	// MinWidthString string `json:"min_width,omitempty"`
	// MinWidthPixel  int    `json:"min_width,omitempty"`

	// center left right
	Align Alignment `json:"align,omitempty"`

	Name                string `json:"name,omitempty"`
	Instance            string `json:"instance,omitempty"`
	Urgent              bool   `json:"urgent,omitempty"`
	Separator           bool   `json:"separator,omitempty"`
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"`
	Markup              string `json:"markup,omitempty"`

	// should allow custom _keyed things
	// Other Custom
}
