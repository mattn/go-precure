# go-precure

golang port of [Acme::PrettyCure](https://github.com/kan/p5-acme-prettycure)

## Usage

```go
package main

import (
	"fmt"
	"github.com/mattn/go-precure"
)

func main() {
	for _, a := range precure.AllStars {
		fmt.Println(a.HumanName())
	}
}
```

## Installation

```
$ go get github.com/mattn/go-precure
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
