package mod

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readInt(fpath string) (int, error) {
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(string(file)))
}

func readFloat(fpath string) (float64, error) {
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(string(file)), 64)
}

func readString(fpath string) (string, error) {
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(file)), nil
}
