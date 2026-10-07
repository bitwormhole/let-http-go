package main

import (
	"os"

	"github.com/bitwormhole/let-http-go/modules/lethttpgo"
	"github.com/starter-go/starter"
)

func main() {

	a := os.Args
	m := lethttpgo.Module()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
