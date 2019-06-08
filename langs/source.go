package langs

import (
	"fmt"
	"path/filepath"
	"strings"
)

// File represents a destination file to be written
// ID is the relative path of the file in its originating directory
// Data is a sequence of hexa string chunks representing the file as a binary resource
type File struct {
	ID   string
	Data [][]string
}

// Source represents a source model bound to a specific language
// Name is the name of the file to be written without extension
// NS is the destination namespace from where resources should be fetched
// Template is the source code template to be used for writing the file
// Extension is the destination file extension (eg .h for cpp lang)
// Files are the files to be written as binary resources accessible within NS
type Source struct {
	Name      string
	NS        string
	Template  string
	Extension string
	Files     []*File
}

// GetFileName returns the definitive file name to be written
func (src Source) GetFileName() string {
	return trimExtension(src.Name) + src.Extension
}

// NewSource generates a source model, bound to a specific language
// with a preallocated files storage
func NewSource(lang, name, ns string) (*Source, error) {
	switch lang {
	case "cpp":
		return NewCppSource(name, ns), nil
	default:
		return nil, fmt.Errorf("unknown language %s", lang)
	}
}

func trimExtension(f string) string {
	return strings.TrimSuffix(f, filepath.Ext(f))
}
