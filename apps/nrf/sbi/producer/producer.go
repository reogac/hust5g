package producer

import (
	"hust5g/apps/nrf/context"
	"hust5g/sbi"
	"net/http"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbiproducer"})
}

type Producer struct {
	backend Backend
}

type Backend interface {
	Context() *context.Context
}

func New(backend Backend) *Producer {
	return &Producer{
		backend: backend,
	}
}

// create groups of service routes to register to http server
func (p *Producer) Groups() []sbi.HttpRouteGroup {
	management := sbi.HttpRouteGroup{
		Path: "mngr",
		Routes: []sbi.HttpRoute{
			sbi.HttpRoute{
				Name:        "Register",
				Method:      http.MethodPost,
				Pattern:     "reg",
				HandlerFunc: p.OnRegister,
			},
			sbi.HttpRoute{
				Name:        "Heartbeat",
				Method:      http.MethodPost,
				Pattern:     "beat/:nfid",
				HandlerFunc: p.OnHeartbeat,
			},
		},
	}

	groups := []sbi.HttpRouteGroup{
		management,
	}
	return groups
}
