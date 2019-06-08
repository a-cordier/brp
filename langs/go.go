package langs

const golang = `
package {{ .NS }}

var data = map[string][]byte {
	{{ range .Files }}
	"{{ .ID }}": {
		{{ range .Data }} {{ Join . ", " }}, 
		{{ end }}
	},
	{{ end }}
}

func Get(file string) []byte {
	d, ok := data[file]
	if !ok {
		return nil
	}
	return d
}
`

// NewGoSource returns a binary data go source model
func NewGoSource(name, ns string) *Source {
	return &Source{
		name,
		ns,
		golang,
		".go",
		[]*File{},
	}
}
