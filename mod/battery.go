package mod

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"os"
	"path"
	"time"

	"github.com/seankhliao/go3status/protocol"
)

type Battery struct {
	Mod

	Battery           string // battery id
	ColorCharge       string // text color when charging
	ColorCritical     string // text color when discharging < ThresholdCritical
	ColorDischarge    string // text color when discharging (normal)
	ColorWarn         string // text color when discharging < ThresholdWarn
	FormatCharge      string // text/template
	FormatDischarge   string // text/template
	ThresholdCritical int    // Use ColorCritical when disharging < ThresholdCritical
	ThresholdWarn     int    // Use ColorWarn when disharging < ThresholdWarn

	filepath string             // base dir of battery
	tmpl     *template.Template // parsed formats
	// Placeholders:
	// {{.Battery}}
	// {{.Percent}}
	// {{.Status}}
}

func DefaultBattery() Module {
	return &Battery{
		Mod: NewMod("battery", 1),

		Battery:           "BAT0",
		ColorCharge:       "#ebcb8b",
		ColorCritical:     "#bf616a",
		ColorDischarge:    "#d8dee9",
		ColorWarn:         "#d08770",
		FormatCharge:      "⚡: {{.Percent}}%",
		FormatDischarge:   "⊙ : {{.Percent}}%",
		ThresholdCritical: 10,
		ThresholdWarn:     25,
	}
}

func (m *Battery) NewBlock(t time.Time) *protocol.Block {
	status, err := readString(path.Join(m.filepath, "status"))
	if err != nil {
		log.Println(err)
	}
	percent, err := readInt(path.Join(m.filepath, "capacity"))
	if err != nil {
		log.Println(err)
	}
	if percent > 100 {
		percent = 100
	}

	temp := "discharge"
	color := m.ColorDischarge
	if status != "Discharging" {
		// charging
		color = m.ColorCharge
		temp = "charge"
	} else if percent <= m.ThresholdCritical {
		color = m.ColorCritical
	} else if percent <= m.ThresholdWarn {
		color = m.ColorWarn
	}

	data := map[string]interface{}{
		"Battery": m.Battery,
		"Percent": percent,
		"Status":  status,
	}
	var b bytes.Buffer
	if err := m.tmpl.ExecuteTemplate(&b, temp, data); err != nil {
		log.Println(err)
	}

	return &protocol.Block{
		FullText: b.String(),
		Color:    color,
		Name:     m.name,
		Instance: m.Instance,
	}
}

func (m *Battery) Start(blocks []*protocol.Block, pos int) error {
	var err error

	m.filepath = path.Join("/sys", "class", "power_supply", m.Battery)
	if _, err := os.Stat(m.filepath); os.IsNotExist(err) {
		return errors.New("battery not found: " + m.Battery)
	}

	m.tmpl, err = template.New("charge").Parse(m.FormatCharge)
	if err != nil {
		return err
	}
	m.tmpl, err = m.tmpl.New("discharge").Parse(m.FormatDischarge)
	if err != nil {
		return err
	}

	m.Mod.Start(blocks, pos, m.NewBlock)
	return nil
}
