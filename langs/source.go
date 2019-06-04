package langs

import (
	"errors"
	"fmt"
)

type File struct {
	ID   string
	Data [][]string
}

type Source struct {
	Name      string
	NS        string
	Template  string
	Extension string
	Files     []*File
}

func (src Source) GetFileName() string {
	return src.Name + src.Extension
}

func NewSource(lang, name, ns string) (*Source, error) {
	switch lang {
	case "cpp":
		return NewCppSource(name, ns), nil
	default:
		return nil, errors.New(fmt.Sprintf("unknown language %s", lang))
	}
}
