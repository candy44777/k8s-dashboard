package namespace

import (
	"github.com/candy44777/k8s-dashboard/protocol"
	"github.com/candy44777/k8s-dashboard/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ctl struct {
	Service *Service
}

func NewCtl() *Ctl {
	return &Ctl{
		Service: NewService(),
	}
}

func (c *Ctl) QueryAll(ctx *gin.Context) {
	req := &RequestByNamespace{}
	ns, err := c.Service.QueryAll(ctx, req)
	if err != nil {
		tools.ResponseFailed(ctx, http.StatusBadRequest, err)
		return
	}
	tools.ResponseSuccess(ctx, http.StatusOK, ns)
}

func (c *Ctl) Get(ctx *gin.Context) {
	name := ctx.Param("name")
	ns, err := c.Service.GetByName(ctx, &RequestByNamespace{Name: name})
	if err != nil {
		tools.ResponseFailed(ctx, http.StatusBadRequest, err)
		return
	}
	tools.ResponseSuccess(ctx, http.StatusOK, ns)
}

func (c *Ctl) Delete(ctx *gin.Context) {
	name := ctx.Param("name")
	ns, err := c.Service.Delete(ctx, &RequestByNamespace{Name: name})
	if err != nil {
		tools.ResponseFailed(ctx, http.StatusBadRequest, err)
		return
	}
	tools.ResponseSuccess(ctx, http.StatusOK, ns)
}

func (c *Ctl) Create(ctx *gin.Context) {
	ns := &Namespace{}
	if err := ctx.ShouldBindJSON(ns); err != nil {
		tools.ResponseFailed(ctx, http.StatusBadRequest, err)
		return
	}

	ns, err := c.Service.Create(ctx, ns)
	if err != nil {
		tools.ResponseFailed(ctx, http.StatusBadRequest, err)
		return
	}

	tools.ResponseSuccess(ctx, http.StatusOK, ns)
}

func (c *Ctl) PathLabel(ctx *gin.Context) {
	req := &RequestByLabel{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		tools.ResponseFailed(ctx, http.StatusBadRequest, err)
		return
	}

	ns, err := c.Service.OverwriteLabel(ctx, req)
	if err != nil {
		tools.ResponseFailed(ctx, http.StatusBadRequest, err)
		return
	}

	tools.ResponseSuccess(ctx, http.StatusOK, ns)
}

func (c *Ctl) Name() string {
	return "namespaceCtl"
}

func (c *Ctl) Build(server *protocol.HttpServer) {
	server.Handle("GET", "/nslist", c.QueryAll)
	server.Handle("GET", "/ns/:name", c.Get)
	server.Handle("DELETE", "/ns/:name", c.Delete)
	server.Handle("POST", "/ns", c.Create)
	server.Handle("PATCH", "/ns", c.PathLabel)
}
