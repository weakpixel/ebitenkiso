package aseprite

type SpriteSheet struct {
	Frames Frames `json:"frames"`
	Meta   Meta   `json:"meta"`
}
type Frame struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}
type SpriteSourceSize struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}
type SourceSize struct {
	W int `json:"w"`
	H int `json:"h"`
}
type FrameMeta struct {
	Frame            Frame            `json:"frame"`
	Rotated          bool             `json:"rotated"`
	Trimmed          bool             `json:"trimmed"`
	SpriteSourceSize SpriteSourceSize `json:"spriteSourceSize"`
	SourceSize       SourceSize       `json:"sourceSize"`
	Duration         int              `json:"duration"`
}

type Frames []FrameMeta

type Size struct {
	W int `json:"w"`
	H int `json:"h"`
}
type FrameTags struct {
	Name      string `json:"name"`
	From      int    `json:"from"`
	To        int    `json:"to"`
	Direction string `json:"direction"`
}
type Layers struct {
	Name      string `json:"name"`
	Opacity   int    `json:"opacity"`
	BlendMode string `json:"blendMode"`
}

type Meta struct {
	App       string      `json:"app"`
	Version   string      `json:"version"`
	Image     string      `json:"image"`
	Format    string      `json:"format"`
	Size      Size        `json:"size"`
	Scale     string      `json:"scale"`
	FrameTags []FrameTags `json:"frameTags"`
	Layers    []Layers    `json:"layers"`
	Slices    []any       `json:"slices"`
}
