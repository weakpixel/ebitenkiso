package aseprite

import (
	"image"
	"time"

	"github.com/weakpixel/ebitenkiso/pkg/res"
	"github.com/weakpixel/ebitenkiso/pkg/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadSprite(resource res.Resource) (*sprites.Sprite, error) {

	sp, err := LoadResource(resource)
	if err != nil {
		return nil, err
	}
	dir := res.Dir(resource)
	srcImg := res.Join(dir, sp.Meta.Image)

	img, err := res.Image(srcImg)
	if err != nil {
		return nil, err
	}
	sheet := ToSpriteSheet(sp, img)
	sprite := sprites.NewSprite(sheet)
	if len(sprite.SpriteSheet().Tags) > 0 {
		sprite.SetAnimation(sprite.SpriteSheet().Tags[0].Name, true)
	}

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
