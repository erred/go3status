package mod

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os/exec"
	"time"

	"github.com/seankhliao/go-i3bar-protocol"
)

type Volume struct {
	Mod

	Control    string // control name (from amixer)
	ColorMute  string
	ColorOn    string
	FormatMute string
	FormatOn   string

	prog      string             // cmd, should be amixer
	progFlags []string           // flags, -M for mapped, get to get status
	tmpl      *template.Template // parsed formats
	// Placeholders:
	// {{.State}}
	// {{.Percent}}
}

func DefaultVolume() Module {
	return &Volume{
		Mod: NewMod("volume", 1),

		Control:    "Master",
		ColorMute:  "#bf616a",
		ColorOn:    "#d8dee9",
		FormatMute: " ✗ : {{.Percent}}% ",
		FormatOn:   " ♫ : {{.Percent}}% ",

		prog:      "amixer",
		progFlags: []string{"-M", "get"},
	}
}

func (m *Volume) NewBlock(t time.Time) *protocol.Block {
	out, err := exec.Command(m.prog, m.progFlags...).Output()
	if err != nil {
		log.Println(err)
	}

	bb := bytes.Split(bytes.TrimSpace(out), []byte("\n"))
	buf := bytes.NewBuffer(bb[len(bb)-1])
	var volume int

	var status, x string
	format := "%s %s %s [%d%%] %s %s"
	_, err = fmt.Fscanf(buf, format, &x, &x, &x, &volume, &x, &status)
	if err != nil {
		log.Println(err)
	}
	status = status[1 : len(status)-1]

	color := m.ColorOn
	if status != "on" {
		color = m.ColorMute
	}

	data := map[string]interface{}{
		"Percent": volume,
		"Status":  status,
	}
	var b bytes.Buffer
	if err := m.tmpl.ExecuteTemplate(&b, status, data); err != nil {
		log.Println(err)
	}

	return &protocol.Block{
		FullText: b.String(),
		Color:    color,
		Name:     m.name,
		Instance: m.Instance,
	}
}

func (m *Volume) Start(blocks []*protocol.Block, pos int) error {
	var err error

	m.progFlags = append(m.progFlags, m.Control)

	m.tmpl, err = template.New("on").Parse(m.FormatOn)
	if err != nil {
		return err
	}
	m.tmpl, err = m.tmpl.New("off").Parse(m.FormatMute)
	if err != nil {
		return err
	}

	m.Mod.Start(blocks, pos, m.NewBlock)
	return nil
}
