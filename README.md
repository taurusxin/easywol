# easywol

Easiest way to send magic packet to wake up a computer in Golang.

## Installation

```shell
go get -u github.com/taurusxin/easywol
```

## Usage

```go
package main

import "github.com/taurusxin/easywol"

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
