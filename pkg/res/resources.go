package res

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Resource struct {
	p      string
	isHttp bool
	fs     fs.FS
}

func Dir(r Resource) Resource {
	if r.fs != nil {
		return Resource{p: path.Dir(r.p), fs: r.fs}
	}
	if r.isHttp {
		u, _ := url.Parse(r.p)
		u.Path = path.Dir(u.Path)
		return Resource{p: u.String(), isHttp: true}
	}
	return Resource{p: filepath.Dir(r.p)}
}

func Join(r Resource, elem ...string) Resource {
	if r.fs != nil {
		parts := append([]string{r.p}, elem...)
		return Resource{p: path.Join(parts...), fs: r.fs}
	}
	if r.isHttp {
		u, _ := url.Parse(r.p)
		parts := append([]string{u.Path}, elem...)
		u.Path = path.Join(parts...)
		return Resource{p: u.String(), isHttp: true}
	}
	parts := append([]string{r.p}, elem...)
	return Resource{p: filepath.Join(parts...)}
}

func ReadAll(r Resource) ([]byte, error) {
	rc, err := Open(r)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll failed: %s", err)
	}
	return data, nil
}

func Open(r Resource) (io.ReadCloser, error) {
	if r.fs != nil {
		f, err := r.fs.Open(r.p)
		if err != nil {
			return nil, fmt.Errorf("fs.Open failed: %s", err)
		}
		return f, nil
	}
	if r.isHttp {
		res, err := http.Get(r.p)
		if err != nil {
			return nil, fmt.Errorf("http.Get failed: %s", err)
		}
		if res.StatusCode != 200 {
			return nil, fmt.Errorf("http.Get: status %s", res.Status)
		}
		return res.Body, nil
	}

	f, err := os.Open(r.p)
	if err != nil {
		return nil, fmt.Errorf("os.Open failed: %s", err)
	}
	return f, nil
}

func Image(r Resource) (*ebiten.Image, error) {
	rc, err := Open(r)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	image, _, err := ebitenutil.NewImageFromReader(rc)
	if err != nil {
		return nil, fmt.Errorf("image.Decode failed: %s", err)
	}
	return image, nil
}

func Shader(r Resource) (*ebiten.Shader, error) {
	data, err := ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ebiten.NewShader(data)
}

func MustParse(s string) Resource {
	r, err := Parse(s)
	if err != nil {
		panic(fmt.Sprintf("res: Parse failed: %s", err))
	}
	return r
}

func Parse(s string) (Resource, error) {
	u, err := url.Parse(s)
	if err != nil {
		return Resource{}, fmt.Errorf("url.Parse failed: %s", err)
	}
	switch u.Scheme {
	case "res", "file":
		fullPath, err := filepath.Abs(path.Join(u.Host, u.Path))
		if err != nil {
			return Resource{}, fmt.Errorf("filepath.Abs failed: %s", err)
		}
		p := filepath.FromSlash(fullPath)
		return Resource{p: p}, nil
	case "http", "https":
		return Resource{p: s, isHttp: true}, nil
	default:
		fullPath, err := filepath.Abs(filepath.FromSlash(s))
		if err != nil {
			return Resource{}, fmt.Errorf("filepath.Abs failed: %s", err)
		}
		return Resource{p: fullPath}, nil
	}
}

func FromFS(fs fs.FS, p string) Resource {
	return Resource{p: p, fs: fs}
}
