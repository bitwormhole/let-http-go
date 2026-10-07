package lethttpgo

import (
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
	"github.com/starter-go/vlog"
)

type Runner struct {
	context *Context
}

func (inst *Runner) Init(cfg *Configuration) error {

	const network = "tcp"
	addr1 := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	fsys := files.FS()

	root := fsys.NewPath(cfg.RootDir)
	addr2, err := net.ResolveTCPAddr(network, addr1)

	if err != nil {
		return err
	}

	ctx := new(Context)
	ctx.Configuration = cfg
	ctx.Root = root
	ctx.Addr = addr2

	inst.context = ctx
	return nil
}

func (inst *Runner) innerCheckRootDir(root afs.Path) error {

	file := root.GetChild("index.html")
	bin, err := file.GetIO().ReadBinary(nil)

	if err == nil {
		path := file.GetPath()
		size := len(bin)
		vlog.Info("Read %d bytes from file [%s]", size, path)
	} else {
		return err
	}

	return nil
}

func (inst *Runner) Run() error {

	var err error
	ctx := inst.context
	cfg := ctx.Configuration
	root := ctx.Root.String()
	addr := ctx.Addr.String()
	wpath := cfg.WebPath
	engine := gin.Default()

	vlog.Info("www.root = %s", root)
	vlog.Info("www.addr = %s", addr)

	err = inst.innerCheckRootDir(ctx.Root)
	if err != nil {
		return err
	}

	engine.Static(wpath, root)

	return http.ListenAndServe(addr, engine)
}
