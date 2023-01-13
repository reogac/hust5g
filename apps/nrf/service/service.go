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
	server   sbi.SbiServer      //running http server
	producer *producer.Producer //handling Sbi requests received at the server
	context  *context.Context   // NRF context, the anchoring point for all data models of NRF
	config   *config.Config     // loaded NRF config
}

func New(conf *config.Config) (nf *NRF, err error) {
	log.Info("Create NRF service")
	nf = &NRF{
		config: conf,
	}

	//create sbi producer
	nf.producer = producer.New(nf)

	//initilize context
	if nf.context, err = context.New(conf); err != nil {
		nf = nil
	}

	//create an http server (and register routes + handlers)
	nf.server = sbi.NewSbiServer(&conf.Sbi, nf.producer.Groups())
	return
}

func (nf *NRF) Context() *context.Context {
	return nf.context
}

func (nf *NRF) Start() (err error) {
	log.Info("Start NRF service")
	//start the http server
	nf.server.Serve()
	return
}

func (nf *NRF) Terminate() {
	fmt.Println("Terminate NRF service")
	//clean context (close database connection if any)
	nf.context.Clean()
	//terminate the http server
	nf.server.Terminate()
}
