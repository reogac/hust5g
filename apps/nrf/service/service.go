package service

import (
	"fmt"
	"hust5g/apps/nrf/config"
	"hust5g/apps/nrf/context"
	"hust5g/apps/nrf/sbi/producer"
	"hust5g/sbi"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type NRF struct {
	server   sbi.SbiServer
	producer *producer.Producer //handling Sbi requests received at the server
	context  *context.Context   // PCF context
	config   *config.Config     // loaded PCF config
}

func New(conf *config.Config) (nf *NRF, err error) {
	log.Info("Create NRF service")
	nf = &NRF{
		config: conf,
	}

	//create sbi producer
	nf.producer = producer.New(nf)
	if nf.context, err = context.New(conf); err != nil {
		nf = nil
	}
	nf.server = sbi.NewSbiServer(&conf.Sbi, nf.producer.Groups())
	return
}

func (nf *NRF) Context() *context.Context {
	return nf.context
}

func (nf *NRF) Start() (err error) {
	log.Info("Start NRF service")
	nf.server.Serve()
	return
}

func (nf *NRF) Terminate() {
	fmt.Println("Terminate NRF service")
	nf.context.Clean()
	nf.server.Terminate()
}
