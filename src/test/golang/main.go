package main

import (
	"os"

	"github.com/bitwormhole/let-http-go/modules/lethttpgo"
	"github.com/starter-go/units"
)

func main() {

	a := os.Args
	m := lethttpgo.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
