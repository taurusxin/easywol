# easywol
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/taurusxin/easywol)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/taurusxin/easywol?style=flat-square)

Easiest way to send magic packet to wake up a computer in Golang.

## Installation

```shell
go get -u github.com/taurusxin/easywol
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/taurusxin/easywol"
)

func main() {
  packet, err := easywol.CreateMagicPacket("B4:2E:99:EF:01:E3")
  if err != nil {
    fmt.Println(err)
    return
  }
  packet.Send("10.0.0.100")           // Send the packet to the specified address
  packet.SendPort("10.0.0.100", "7")  // Send the packet to the specified address and port
}
```

## License

MIT
