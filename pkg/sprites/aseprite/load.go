package aseprite

import (
	"encoding/json"
	"io"
	"os"
)

func Load(in io.Reader) (*SpriteSheet, error) {
	dec := json.NewDecoder(in)
	sheet := &SpriteSheet{}
	err := dec.Decode(sheet)
	if err != nil {
		return nil, err
	}
	return sheet, nil
}

func LoadFile(filename string) (*SpriteSheet, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Load(f)
}
