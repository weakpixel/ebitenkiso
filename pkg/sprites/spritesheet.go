package sprites

type Tag struct {
	Name string
	From int
	To   int
}

type SpriteSheet struct {
	Tags   []Tag
	Frames []Frame
}

func (s *SpriteSheet) FrameSize() (w, h int) {
	if len(s.Frames) > 0 {
		return s.Frames[0].Width, s.Frames[0].Height
	}
	return 0, 0
}

func (s *SpriteSheet) Add(tag string, frames []Frame) {
	from := len(s.Frames)
	t := Tag{
		Name: tag,
		From: from,
		To:   from + len(frames) - 1,
	}
	s.Frames = append(s.Frames, frames...)
	s.Tags = append(s.Tags, t)
}

func (s *SpriteSheet) TagByName(tag string) *Tag {
	for _, t := range s.Tags {
		if t.Name == tag {
			return &t
		}
	}
	return nil
}

func (s *SpriteSheet) Animation(tag string) *Animation {
	t := s.TagByName(tag)
	if t == nil {
		panic("animation cannot find tag: " + tag)
	}
	frames := s.Frames[t.From : t.To+1]
	return &Animation{
		Frames: frames,
		Name:   t.Name,
	}
}
