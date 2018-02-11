package main

import (
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/seankhliao/go3status/protocol"
	"github.com/seankhliao/go3status/status"
)

func main() {
	conf, err := NewConfig("default.toml")
	if err != nil {
		log.Fatal(err)
	}

	s := status.NewStatus(os.Stdout, protocol.MinimalHeader())
	if err := s.Begin(); err != nil {
		log.Fatal(err)
	}

	ts, bs := conf.StartBlocks()
	s.NewBlocks(bs)

	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		for _, s := range ts {
			s <- t
		}
		s.Next()
	}
}

type Module interface {
	Rename(name, instance string)
	Start() (chan time.Time, chan *protocol.Block)
}

type Config struct {
	// colors = true
	// color_good = "#81a2be"
	// color_degraded = "#b294bb"
	// color_bad = "#cc6666"
	// interval = 1

	// order += "wifi"
	ColorGood     string
	ColorDegraded string
	ColorBad      string
	Interval      int

	RawOpts []RawOpt `toml:"block"`
	Blocks  []Module
}

type RawOpt struct {
	Name     string
	Instance string
	Options  toml.Primitive
}

func NewConfig(tom string) (*Config, error) {
	var c Config
	m, err := toml.DecodeFile(tom, &c)
	if err != nil {
		return &c, err
	}

	for _, opts := range c.RawOpts {
		var mod Module
		switch opts.Name {
		default:
			return &c, err
		case "static":
			mod = &ModStatic{}
		}
		if err := m.PrimitiveDecode(opts.Options, mod); err != nil {
			return &c, err
		}

		mod.Rename(opts.Name, opts.Instance)
		c.Blocks = append(c.Blocks, mod)
	}
	return &c, nil
}

func (c *Config) StartBlocks() ([]chan time.Time, []chan *protocol.Block) {
	var ts []chan time.Time
	var bs []chan *protocol.Block
	for _, block := range c.Blocks {
		t, b := block.Start()
		ts = append(ts, t)
		bs = append(bs, b)
	}
	return ts, bs
}
