package main

import (
	"errors"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/seankhliao/go3status/mod"
	"github.com/seankhliao/go3status/protocol"
)

var ModuleNames = map[string]func(string, string) mod.Module{
	"static": mod.NewStatic,
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

	// RawOpts []RawOpt `toml:"block"`
	RawOpts []struct {
		Name     string
		Instance string
		Options  toml.Primitive
	} `toml:"block"`
	Blocks []mod.Module
}

func ParseConfig(tom string) (*Config, error) {
	var c Config
	m, err := toml.DecodeFile(tom, &c)
	if err != nil {
		return &c, err
	}

	for _, opts := range c.RawOpts {
		if _, ok := ModuleNames[opts.Name]; !ok {
			return &c, errors.New("name not found")
		}

		mod := ModuleNames[opts.Name](opts.Instance)
		if err := m.PrimitiveDecode(opts.Options, mod); err != nil {
			return &c, err
		}
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
