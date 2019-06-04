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

	char* getResource(const char* name) {
		auto it = data.find(name);
		return it == data.end() ? nullptr : it->second.data();
	}
}
`

func NewCppSource(name, ns string) *Source {
	return &Source{
		name,
		ns,
		cpp,
		".h",
		[]*File{},
	}
}
