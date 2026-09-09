goldmark-github-slugger
=========================

[![GoDev][godev-image]][godev-url]

[godev-image]: https://pkg.go.dev/badge/github.com/yuin/goldmark-github-slugger
[godev-url]: https://pkg.go.dev/github.com/yuin/goldmark-github-slugger

goldmark-github-slugger is a `parser.IDGenerator` that generates heading IDs in the same way as GitHub does.

## Compatiblity
`github.com/yuin/goldmark-github-slugger` is compatible with `goldmark/v2`.

## Installation

```sh
go get github.com/yuin/goldmark-github-slugger
```

## Usage

```go
import (
    "bytes"
    "fmt"

    "github.com/yuin/goldmark-github-slugger"
    "github.com/yuin/goldmark/v2/parser"
    "github.com/yuin/goldmark/v2/renderer/html"
)

func main() {
    p := parser.New(parser.WithExtensions(slugger.Parser), parser.WithAutoHeadingID())
    r := html.New()
}
```

## License
MIT

## Author

Yusuke Inuzuka
