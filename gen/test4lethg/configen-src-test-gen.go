package test4lethg
import (
    p3da1d72e8 "github.com/bitwormhole/let-http-go/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type p3da1d72e8.Example4t in package:github.com/bitwormhole/let-http-go/src/test/golang/testcom
//
// id:com-3da1d72e85ff5462-testcom-Example4t
// class:
// alias:
// scope:singleton
//
type p3da1d72e85_testcom_Example4t struct {
}

func (inst* p3da1d72e85_testcom_Example4t) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3da1d72e85ff5462-testcom-Example4t"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3da1d72e85_testcom_Example4t) new() any {
    return &p3da1d72e8.Example4t{}
}

func (inst* p3da1d72e85_testcom_Example4t) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3da1d72e8.Example4t)
	nop(ie, com)

	


    return nil
}


