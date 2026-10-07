package main4lethg
import (
    p51ef1d3e7 "github.com/bitwormhole/let-http-go/app/boot"
     "github.com/starter-go/application"
)

// type p51ef1d3e7.Bootstrap in package:github.com/bitwormhole/let-http-go/app/boot
//
// id:com-51ef1d3e7a0d1347-boot-Bootstrap
// class:
// alias:
// scope:singleton
//
type p51ef1d3e7a_boot_Bootstrap struct {
}

func (inst* p51ef1d3e7a_boot_Bootstrap) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-51ef1d3e7a0d1347-boot-Bootstrap"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p51ef1d3e7a_boot_Bootstrap) new() any {
    return &p51ef1d3e7.Bootstrap{}
}

func (inst* p51ef1d3e7a_boot_Bootstrap) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p51ef1d3e7.Bootstrap)
	nop(ie, com)

	
    com.Root = inst.getRoot(ie)
    com.Host = inst.getHost(ie)
    com.Path = inst.getPath(ie)
    com.Port = inst.getPort(ie)


    return nil
}


func (inst*p51ef1d3e7a_boot_Bootstrap) getRoot(ie application.InjectionExt)string{
    return ie.GetString("${let.http-go.root}")
}


func (inst*p51ef1d3e7a_boot_Bootstrap) getHost(ie application.InjectionExt)string{
    return ie.GetString("${let.http-go.host}")
}


func (inst*p51ef1d3e7a_boot_Bootstrap) getPath(ie application.InjectionExt)string{
    return ie.GetString("${let.http-go.webpath}")
}


func (inst*p51ef1d3e7a_boot_Bootstrap) getPort(ie application.InjectionExt)int{
    return ie.GetInt("${let.http-go.port}")
}


