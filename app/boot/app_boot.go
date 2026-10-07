package boot

import (
	lethttpgo "github.com/bitwormhole/let-http-go"
	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type Bootstrap struct {

	//starter:component

	Root string //starter:inject("${let.http-go.root}")
	Host string //starter:inject("${let.http-go.host}")
	Path string //starter:inject("${let.http-go.webpath}")

	Port int //starter:inject("${let.http-go.port}")

}

func (inst *Bootstrap) _impl() application.Lifecycle {
	return inst
}

func (inst *Bootstrap) Life() *application.Life {
	l := new(application.Life)

	l.OnCreate = inst.onCreate
	l.OnStart = inst.onStart
	l.OnLoop = inst.onLoop
	l.OnStop = inst.onStop
	l.OnDestroy = inst.onDestroy

	return l
}

func (inst *Bootstrap) onCreate() error {
	vlog.Debug("Bootstrap:onCreate()")
	return nil
}

func (inst *Bootstrap) onStart() error {
	vlog.Debug("Bootstrap:onStart()")
	return nil
}

func (inst *Bootstrap) onLoop() error {
	// vlog.Debug("Bootstrap:onLoop()")

	var err error
	cfg := inst.innerGetConfig()
	r1 := new(lethttpgo.Runner)

	err = r1.Init(cfg)
	if err != nil {
		return err
	}

	err = r1.Run()
	if err != nil {
		return err
	}

	return nil
}

func (inst *Bootstrap) onStop() error {
	vlog.Debug("Bootstrap:onStop()")
	return nil
}

func (inst *Bootstrap) onDestroy() error {
	vlog.Debug("Bootstrap:onDestroy()")
	return nil
}

func (inst *Bootstrap) innerGetConfig() *lethttpgo.Configuration {

	cfg := new(lethttpgo.Configuration)

	cfg.Host = inst.Host
	cfg.Port = inst.Port
	cfg.RootDir = inst.Root
	cfg.WebPath = inst.Path

	return cfg
}

////////////////////////////////////////////////////////////////////////////////
