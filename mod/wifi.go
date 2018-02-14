package mod

import (
	"bytes"
	"html/template"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/seankhliao/go-i3bar-protocol"
)

type Wifi struct {
	Mod

	Interface          string
	ColorConnected     string
	ColorDisconnected  string
	FormatConnected    string
	FormatDisconnected string

	prog      string             // cmd, should be amixer
	progFlags []string           // flags, -M for mapped, get to get status
	tmpl      *template.Template // parsed formats
	// Placeholders:
	// {{.State}}
	// {{.Percent}}
}

func DefaultWifi() Module {
	return &Wifi{
		Mod: NewMod("volume", 1),

		Interface:          "wlp58s0",
		ColorConnected:     "#d8dee9",
		ColorDisconnected:  "#bf616a",
		FormatConnected:    " {{.Ssid}} : {{.Ip}} ",
		FormatDisconnected: " ✗ : No Connection ",

		prog:      "wpa_cli",
		progFlags: []string{"status", "-i"},
	}
}

func (m *Wifi) NewBlock(t time.Time) *protocol.Block {
	out, err := exec.Command(m.prog, m.progFlags...).Output()
	if err != nil {
		log.Println(err)
	}

	bb := bytes.Split(bytes.TrimSpace(out), []byte("\n"))
	rawData := map[string]string{}
	for _, b := range bb {
		kv := bytes.Split(bytes.TrimSpace(b), []byte("="))
		rawData[string(kv[0])] = string(kv[1])
	}

	status := strings.ToLower(rawData["wpa_state"])
	color := m.ColorDisconnected
	if status != "disconnected" {
		color = m.ColorConnected
		status = "connected"
	}

	data := map[string]interface{}{
		"Status": rawData["wpa_state"],
		"Ip":     rawData["ip_address"],
		"Ssid":   rawData["ssid"],
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

func (m *Wifi) Start(blocks []*protocol.Block, pos int) error {
	var err error

	m.progFlags = append(m.progFlags, m.Interface)

	m.tmpl, err = template.New("connected").Parse(m.FormatConnected)
	if err != nil {
		return err
	}
	m.tmpl, err = m.tmpl.New("disconnected").Parse(m.FormatDisconnected)
	if err != nil {
		return err
	}

	m.Mod.Start(blocks, pos, m.NewBlock)
	return nil
}
