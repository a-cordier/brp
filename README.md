# brp

## Binary Resources Packager

brp (binary resources packager) packages files as binary resources availables in your project

brp was originally written in cpp as [bpk](https://github.com/a-cordier/bpk) and intended to be used for cpp projects.

The current project goal is to take advantage of go templates and the [cobra](https://github.com/spf13/cobra) library to:

  - make bpk easily extensible to other languages
  - provide a more intuitive command line tool

If you want pure cpp integration with cmake, you may still want to use brp

## Current Language Support

  - [cpp](#cpp)
  - [golang](#golang)
  
## Requirements

brp is packaged as a go module, so you will need golang v1.12+

## Build from source

```sh
git clone https://github.com/a-cordier/brp.git
cd brp
go build -o brp main.go
```

## Package your files

```sh
brp generate <resource_directory> -o resources.h -l cpp -n resources
```

  - Ommiting the -o flag will result in writing to stdout
  - Ommiting the -l flag will result in using cpp language
  - Ommiting the -n flag will result in using the <resource_directory> name converted to a legal camelCased string (no leading numeric - no special character) as a namespace

## Examples

### cpp

Assuming the following resource directory:

```
resources
└── svg
    ├── next.svg
    ├── pause.svg
    ├── play.svg
    ├── previous.svg
    └── stop.svg
```

Running

```sh
brp generate resources -o resources.h -n resources -l cpp
```

Will generate the following `resources.h` file:

```cpp
#pragma once

#include <iostream>
#include <vector>
#include <map>
#include <utility>

namespace resources {

	namespace {

		std::map<std::string, std::vector<char> > data = {
			{ "svg/previous.svg", { /* Data chunks */ } },
			{ "svg/pause.svg", { /* Data chunks */ } },
			{ "svg/play.svg", { /* Data chunks */ } },
			{ "svg/next.svg", { /* Data chunks */ } },
		};
	}

	inline char* get(const char* name) {
		auto it = data.find(name);
		return it == data.end() ? nullptr : it->second.data();
	}

	inline std::vector::size_type size(const char* name) {
		auto it = data.find(name);
		return it == data.end() ? 0 : it->second.size();
	}
}
```

Resources being accessed the following way

```cpp
#include "resources.h"

auto data = resources::get("svg/play.svg")
```

### golang

Assuming the following resource directory:

```
resources
└── svg
    ├── next.svg
    ├── pause.svg
    ├── play.svg
    ├── previous.svg
    └── stop.svg
```

Running

```sh
brp generate resources -o resources.go -n resources -l go
```

Will generate the following `resources.go` file:

```go

package resources

var data = map[string][]byte {
	
	"svg/next.svg": {
		 // Data chunks		
	},
	
	"svg/pause.svg": {
		// Data chunks 
	},
	
	"svg/play.svg": {
		// Data chunks
	},
	
	"svg/previous.svg": {
		 // Data chunks
	},
	
	"svg/stop.svg": {
		 // Data chunks
	},
	
}

func Get(file string) []byte {
	d, ok := data[file]
	if !ok {
		return nil
	}
	return d
}
```

Resources being accessed the following way

```cpp
import "github.com/<user>/<project>/resources"

data := resources.Get("svg/play.svg")
```
