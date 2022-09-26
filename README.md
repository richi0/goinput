# Goinput

Package to get input from stdin or non-flag command-line arguments.

## Installation

`go get github.com/richi0/goinput`

## Usage example

Simple filter that takes input from stdin or command-line and returns the uppercase version in stdout.

```go
package main

import (
	"fmt"
	"strings"

	input "github.com/richi0/goinput"
)

func main() {
	input := input.GetInput()
	fmt.Printf("%s", strings.ToUpper(input))
}
```
