package ginx

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	IP   string
	Port int
	app  *gin.Engine
}

func (gs *GinServer) SetDefaults() {
	if gs.Port == 0 {
		gs.Port = 80
	}
}

func (gs *GinServer) Run() {
	addr := gs.ListenAddr()
	gs.app.Run(addr)
}

func (gs *GinServer) ListenAddr() string {
	return fmt.Sprintf("%s:%d", gs.IP, gs.Port)
}

type GinEngineFunc func(e *gin.Engine)

func (gs *GinServer) RegisterEngine(fn GinEngineFunc) {
	fn(gs.app)
}
