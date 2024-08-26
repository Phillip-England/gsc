package gsc

import "github.com/Phillip-England/ffh"

type GscFile struct {
	Path       string
	Components []Component
	Content    string
}

func NewGscFile(path string) (*GscFile, error) {
	f := &GscFile{
		Path:       path,
		Components: make([]Component, 0),
		Content:    "",
	}
	content, err := f.Read()
	if err != nil {
		return nil, err
	}
	f.Content = content
	return f, nil
}

func (f *GscFile) Read() (string, error) {
	content, err := ffh.ReadFile(f.Path)
	if err != nil {
		return "", err
	}
	return content, nil
}

type Component struct {
	Params map[string]string
}
