# garm
A simple ARM converter for Golang

# Installation

```
go get github.com/HIDE810/garm
```

# Usage

```Go
package main

import (
	"fmt"
	"log"

	"github.com/HIDE810/garm"
)

func main() {
	hex, err := garm.ArmToHex("mov r0, r0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex)
}
```

# API
It consumes data from [armconverter](https://armconverter.com/).

# License
MIT
