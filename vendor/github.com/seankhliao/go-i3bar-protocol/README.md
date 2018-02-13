# go-i3bar-protocol [![GoDoc](https://godoc.org/github.com/seankhliao/go-i3bar-protocol?status.svg)](https://godoc.org/github.com/seankhliao/go-i3bar-protocol) [![Build Status](https://img.shields.io/travis/seankhliao/go-i3bar-protocol.svg?style=flat-square)](https://travis-ci.org/seankhliao/go-i3bar-protocol) [![Go Report Card](https://goreportcard.com/badge/github.com/seankhliao/go-i3bar-protocol)](https://goreportcard.com/report/github.com/seankhliao/go-i3bar-protocol)

Go types for use with [i3bar](https://i3wm.org/docs/i3bar-protocol.html)

## Install

```bash
go get github.com/seankhliao/go-i3bar-protocol
```

## Example

```go
package main

import (
        "github.com/seankhliao/go-i3bar-protocol"
        "encoding/json"
        "fmt"
        "os"
)

func main() {
        e := protocol.NewEncoder(os.Stdout)
        e.Encode(protocol.MinimalHeader())
        e.BeginArray()

        for {
                var blocks []protocol.Block
                // fill blocks
                e.Encode(blocks)
        }
}
```

## License

The MIT License (MIT) - see [`LICENSE`](LICENSE) for more details
