package aseprite

import (
	"image"
	"path/filepath"
	"time"

	"ebitenkiso/pkg/sprites"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadSprite(animFile string, assetBase string) (*sprites.Sprite, error) {
	if assetBase == "" {
		assetBase = filepath.Dir(animFile)
	}

	sp, err := LoadFile(animFile)
	if err != nil {
		return nil, err
	}

	srcImg := filepath.Join(assetBase, sp.Meta.Image)
	img, _, err := ebitenutil.NewImageFromFile(srcImg)
	if err != nil {
		return nil, err
	}
	sheet := ToSpriteSheet(sp, img)
	sprite := sprites.NewSprite(sheet)
	sprite.Src = srcImg
	return sprite, nil
}

func ToSpriteSheet(sheet *SpriteSheet, img *ebiten.Image) *sprites.SpriteSheet {
	tags := make([]sprites.Tag, len(sheet.Meta.FrameTags))
	for idx, t := range sheet.Meta.FrameTags {
		tags[idx] = sprites.Tag{
			Name: t.Name,
			From: t.From,
			To:   t.To,
		}
	}
	frames := make([]sprites.Frame, len(sheet.Frames))
	for idx, f := range sheet.Frames {
		frames[idx] = sprites.Frame{
			Duration: time.Duration(f.Duration) * time.Millisecond,
			Image:    img.SubImage(image.Rect(f.Frame.X, f.Frame.Y, f.Frame.X+f.Frame.W, f.Frame.Y+f.Frame.H)).(*ebiten.Image),
			Width:    f.Frame.W,
			Height:   f.Frame.H,
		}
	}

	return &sprites.SpriteSheet{
		Frames: frames,
		Tags:   tags,
	}
}
