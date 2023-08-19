package main

import (
	"runtime"

	"github.com/fish895623/ticker/hello"
)

func main() {
	runtime.GOMAXPROCS(2)
	clone.Hello()
}
