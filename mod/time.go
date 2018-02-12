package mod

import (
	"bytes"
	"html/template"
	"log"
	"time"

	"github.com/seankhliao/go3status/protocol"
)

type Time struct {
	Mod

	// Specific

	TimeFormat string
	TimeZone   string
	Format     string

	location *time.Location
	tmpl     *template.Template
}

func NewTime() Module {
	return &Time{
		Mod: NewMod("time", 1),

		TimeFormat: time.ANSIC,
		TimeZone:   "Local",
		Format:     "{{.Timezone}}: {{.Time}}",
	}
}

func (m *Time) NewBlock(t time.Time) *protocol.Block {
	data := map[string]interface{}{
		"Time":     t.In(m.location).Format(m.TimeFormat),
		"Timezone": m.location.String(),
	}

	var b bytes.Buffer
	err := m.tmpl.Execute(&b, data)
	if err != nil {
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

	m.tmpl, err = template.New("t").Parse(m.Format)
	if err != nil {
		return err
	}

	m.Mod.Start(blocks, pos, m.NewBlock)
	return nil
}
