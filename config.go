package main

import (
	"errors"

	"github.com/pelletier/go-toml"
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

	tree, err := toml.LoadFile(fpath)
	if err != nil {
		return blocks, err
	}

	for _, raw := range tree.Get("conf").([]*toml.Tree) {
		key := raw.Keys()[0]
		newBlockFunc, ok := ModuleNames[key]
		if !ok {
			return blocks, errors.New("module not found: " + key)
		}

		block := newBlockFunc()
		if err = raw.Get(key).(*toml.Tree).Unmarshal(block); err != nil {
			return blocks, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}
