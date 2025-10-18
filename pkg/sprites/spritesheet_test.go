package sprites

import (
	"testing"
	"time"
)

func TestSpritesheet(t *testing.T) {
	s := &SpriteSheet{}
	s.Add("run", []Frame{
		{Duration: time.Millisecond * 10},
		{Duration: time.Millisecond * 10},
		{Duration: time.Millisecond * 10},
	})
	s.Add("jump", []Frame{
		{Duration: time.Millisecond * 20},
	})

	anim := s.Animation("run")
	if len(anim.Frames) != 3 {
		t.Errorf("run animation must have 3 frames but has %d", len(anim.Frames))
	}

	for _, f := range anim.Frames {
		if f.Duration != time.Millisecond*10 {
			t.Error("jump animation contains wrong frames")
		}
	}
	for range anim.Frames {
		anim.Update(time.Millisecond * 11)
	}

	if anim.frameIndex != len(anim.Frames)-1 {
		t.Errorf("run animation must be on last index: %d but is %d", len(anim.Frames)-1, anim.frameIndex)
	}

	anim = s.Animation("jump")
	if len(anim.Frames) != 1 {
		t.Errorf("jump animation must have 1 frame but has %d", len(anim.Frames))

	}
	if anim.Frames[0].Duration != time.Millisecond*20 {
		t.Error("jump animation contains wrong frames")
	}

	anim.Update(time.Millisecond * 30)
	if anim.frameIndex != 0 {
		t.Errorf("jump has only one frame frameIndex must be zero but is: %d", anim.frameIndex)
	}
	anim.Update(time.Millisecond * 30)
	if anim.frameIndex != 0 {
		t.Errorf("jump has only one frame frameIndex must be zero but is: %d", anim.frameIndex)
	}

}

func TestAnimLoop(t *testing.T) {
	s := &SpriteSheet{}
	s.Add("run", []Frame{
		{Duration: time.Millisecond * 10},
		{Duration: time.Millisecond * 10},
		{Duration: time.Millisecond * 10},
	})

	anim := s.Animation("run")
	anim.Loop = true

	for range anim.Frames {
		anim.Update(time.Millisecond * 11)
	}
	for range anim.Frames {
		anim.Update(time.Millisecond * 11)
	}
	for range anim.Frames {
		anim.Update(time.Millisecond * 11)
	}

	anim.Reset()

}
