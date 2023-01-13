package main

import (
	"hust5g/apps/nrf/config"
	"hust5g/apps/nrf/service"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Usage: "Load configuration from `FILE`",
	},
	cli.StringFlag{
		Name:  "log, l",
		Usage: "Output logs to `FILE`",
	},
}

var nf *service.NRF

func main() {
	log.Println("Hello beautiful world")

	app := cli.NewApp()
	app.Name = "nrf"
	app.Usage = "5G simple NRF"
	app.Action = action
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		//log
		log.Fatal("Fail to start application", err)
	} else {
		quit := make(chan struct{})
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-sigch
			if nf != nil {
				nf.Terminate()
			}
			log.Info("Received a kill signal")
			quit <- struct{}{}
		}()
		<-quit
		log.Info("Good bye the world")
	}
}

func action(c *cli.Context) (err error) {
	log.SetLevel(log.InfoLevel)
	//read config
	var cfg config.Config
	filename := c.String("config")
	if cfg, err = config.Load(filename); err != nil {
		log.Errorf("Fail to parse configuration", err)
		return
	}

	//create the NRF
	if nf, err = service.New(&cfg); err != nil {
		log.Errorf("Fail to create NRF", err)
		return
	}

	if err = nf.Start(); err != nil {
		log.Errorf("Fail to start NRF", err)
		return
	}

	return
}
