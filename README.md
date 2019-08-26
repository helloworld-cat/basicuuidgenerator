# Basic UUID Generator

This package generate UUID with crypto lib [RFC 4122](http://www.ietf.org/rfc/rfc4122.txt).

## Installation

Use the `go` tool:

	$ go get github.com/pagedegeek/basicuuidgenerator

## Usage

```
// ...

import (
    "github.com/pagedegeek/basicuuidgenerator"
)

// ...

uuid, err := basicuuidgenerator.NewV4()

// ...
```

