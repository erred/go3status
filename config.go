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

// type General struct {
// 	ColorGood     string
// 	ColorDegraded string
// 	ColorBad      string
// 	// Interval      int
// }

func ParseConfig(fpath string) ([]mod.Module, error) {
	var rawConf map[string]toml.Primitive
	var blocks []mod.Module

	m, err := toml.DecodeFile(fpath, &rawConf)
	if err != nil {
		return blocks, err
	}

	for name, primitive := range rawConf {
		newBlockFunc, ok := ModuleNames[name]
		if !ok {
			return blocks, errors.New("module not found: " + name)
		}

		block := newBlockFunc()
		if err := m.PrimitiveDecode(primitive, block); err != nil {
			return blocks, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}
