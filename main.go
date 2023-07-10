package main

import (
	"os"

	"github.com/florianrusch/gitsynchro/cmd"
)

func main() {
	_ = cmd.Execute(os.Args[1:])
}
