package lethttpgo

import (
	lethttpgo "github.com/bitwormhole/let-http-go"
	"github.com/bitwormhole/let-http-go/gen/main4lethg"
	"github.com/bitwormhole/let-http-go/gen/test4lethg"
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
)

////////////////////////////////////////////////////////////////////////////////

func Module() application.Module {
	return ModuleForMain()
}

////////////////////////////////////////////////////////////////////////////////

func ModuleForMain() application.Module {

	mb := lethttpgo.BuildModuleForMain()

	mb.Components(main4lethg.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := lethttpgo.BuildModuleForTest()

	mb.Components(test4lethg.ExportComponents)

	mb.Depend(Module())
	mb.Depend(units.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
