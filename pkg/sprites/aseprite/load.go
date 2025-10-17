package aseprite

import (
	"encoding/json"
	"io"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/weakpixel/ebitenkiso/pkg/res"
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
func LoadResource(resource res.Resource) (*SpriteSheet, error) {
	r, err := res.Open(resource)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return Load(r)
}

func LoadFile(filename string) (*SpriteSheet, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Load(f)
}
