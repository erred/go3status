package main

import (
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/seankhliao/go3status/mod"
)

var ModuleNames = map[string]func() mod.Module{
	"static": mod.NewStatic,
	"time":   mod.NewTime,
}

type Config struct {
	ColorGood     string
	ColorDegraded string
	ColorBad      string
	// Interval      int

	RawOpts []struct {
		Name    string
		Options toml.Primitive
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

		mod := ModuleNames[opts.Name]()
		if err := m.PrimitiveDecode(opts.Options, mod); err != nil {
			return &c, err
		}
		c.Blocks = append(c.Blocks, mod)
	}
	return &c, nil
}
