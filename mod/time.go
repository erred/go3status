package mod

import (
	"bytes"
	"html/template"
	"log"
	"time"

	"github.com/seankhliao/go-i3bar-protocol"
)

type Time struct {
	Mod

	Format     string // text/template
	TimeFormat string // Output format of time, Go reference time
	TimeZone   string // Timezone or "Local"

	location *time.Location     // parsed TimeZone
	tmpl     *template.Template // parsed Format
	// Placeholders:
	// {{.Time}}
	// {{.Timezone}}
}

func DefaultTime() Module {
	return &Time{
		Mod: NewMod("time", 1),

		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Format:     " {{.Time}} ",
	}
}

func (m *Time) NewBlock(t time.Time) *protocol.Block {
	data := map[string]interface{}{
		"Time":     t.In(m.location).Format(m.TimeFormat),
		"Timezone": m.location.String(),
	}

	var b bytes.Buffer
	if err := m.tmpl.Execute(&b, data); err != nil {
		log.Println(err)
	}

	return &protocol.Block{
		FullText: b.String(),
		Name:     m.name,
		Instance: m.Instance,
	}
}

func (m *Time) Start(blocks []*protocol.Block, pos int) error {
	var err error

	m.location, err = time.LoadLocation(m.TimeZone)
	if err != nil {
		return err
	}

	m.tmpl, err = template.New("").Parse(m.Format)
	if err != nil {
		return err
	}

	m.Mod.Start(blocks, pos, m.NewBlock)
	return nil
}
