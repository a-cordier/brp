package langs

const cpp = `
#include <iostream>
#include <map>
#include <utility>
#include <vector>

namespace {{ .NS }} {

	namespace {

		std::map<std::string, std::vector<char> > data = {
			{{ range .Files }}
			{
				"{{ .ID }}", {
					{{ range .Data }} {{ Join . ", " }}, 
					{{ end }}
				}
			},
			{{ end }}
		};
	}

	char* get(const char* name) {
		auto it = data.find(name);
		return it == data.end() ? nullptr : it->second.data();
	}
}
`

// NewCppSource returns a binary data cpp source model
func NewCppSource(name, ns string) *Source {
	return &Source{
		name,
		ns,
		cpp,
		".h",
		[]*File{},
	}
}
