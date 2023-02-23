package user

import "github.com/monta-k/go-openapi-gen-example/internal/user/handler"

type Dependency struct {
	pingHandler *handler.PingHandler
}

func (d *Dependency) Inject() {
	d.pingHandler = handler.NewPingHandler()
}
