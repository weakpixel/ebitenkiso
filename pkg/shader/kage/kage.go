package kage

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed *.kage partials/*.kage
	content embed.FS

	shaders = map[string]*ebiten.Shader{}
)

func init() {
	l := &loader{content, map[string]*shader{}}
	list, err := content.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, f := range list {
		if !f.IsDir() {
			key := toKey(f.Name())
			shaders[key], err = l.Load(f.Name())
			if err != nil {
				panic(err)
			}
		}
	}
}

func toKey(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}

var importRegex = regexp.MustCompile("//import:partial (.*)")

type loader struct {
	content fs.FS
	cache   map[string]*shader
}

type shader struct {
	shader *ebiten.Shader
	mod    time.Time
}

func (l *loader) Load(name string) (*ebiten.Shader, error) {
	f, err := l.content.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	mod := stat.ModTime()
	if c, ok := l.cache[name]; ok && c != nil {
		if c.mod == mod {
			return c.shader, nil
		}
	}

	defer f.Close()
	raw, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	matches := importRegex.FindAllSubmatch(raw, -1)
	result := &strings.Builder{}
	result.Write(raw)
	for idx := range matches {
		p := string(matches[idx][1])
		p = strings.TrimSpace(p)
		f, err := l.content.Open(p)
		if err != nil {
			return nil, fmt.Errorf("loading shader %q failed, cannot open partial %q Error: %w", name, p, err)
		}
		raw, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("loading shader %q failed, cannot read partial %q Error: %w", name, p, err)
		}
		f.Close()
		result.WriteString("\n// ============\n")
		result.WriteString("// Import: " + p)
		result.WriteString("\n// ============\n")
		result.Write(raw)
		result.WriteString("\n")
	}
	shaderStr := result.String()
	s, err := ebiten.NewShader([]byte(shaderStr))
	if err != nil {
		return nil, wrapCompileError(name, shaderStr, err)
	}
	l.cache[name] = &shader{
		mod:    mod,
		shader: s,
	}
	return s, nil
}

func wrapCompileError(shaderName string, raw string, err error) error {
	s := bufio.NewScanner(strings.NewReader(raw))
	line := 0
	b := &strings.Builder{}
	for s.Scan() {
		line++
		b.WriteString(fmt.Sprintf("%d %s\n", line, s.Text()))
	}
	return &CompileError{
		shaderName: shaderName,
		code:       b.String(),
		err:        err,
	}
}

type CompileError struct {
	err        error
	shaderName string
	code       string
}

func (err *CompileError) Error() string {
	return fmt.Sprintf("compile failed, shader %q. \n%s\nError: %s", err.shaderName, err.code, err.err)
}
