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
	var blocks []mod.Module
	var base map[string][]map[string]toml.Primitive
	meta, err := toml.DecodeFile(fpath, &base)
	if err != nil {
		return blocks, err
	}

	for _, raw := range base["conf"] {
		var key string
		for k := range raw {
			key = k
		}

		newBlockFunc, ok := ModuleNames[key]
		if !ok {
			return blocks, errors.New("module not found: " + key)
		}

		block := newBlockFunc()
		if err = meta.PrimitiveDecode(raw[key], block); err != nil {
			return blocks, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}
