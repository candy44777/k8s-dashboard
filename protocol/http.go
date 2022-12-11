package protocol

import (
	"github.com/gin-gonic/gin"
)

type HttpRegister interface {
	Name() string
	Build(server *HttpServer)
}

type HttpServer struct {
	*gin.Engine
	group *gin.RouterGroup
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		Engine: gin.New(),
	}
}

func (h *HttpServer) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *HttpServer {
	h.group.Handle(httpMethod, relativePath, handlers...)
	return h
}

func (h *HttpServer) Mount(group string, register ...HttpRegister) *HttpServer {
	h.group = h.Group(group)
	for _, m := range register {
		m.Build(h)
	}
	return h
}
