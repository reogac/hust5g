package context

import (
	"hust5g/apps/nrf/config"
	"hust5g/apps/nrf/db"
	"hust5g/sbi/models"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "context"})
}

type Context struct {
	conf      *config.Config
	profileDb *db.Database
}

func New(conf *config.Config) (ctx *Context, err error) {
	ctx = &Context{
		conf: conf,
	}
	if ctx.profileDb, err = db.Load(conf.DbName); err != nil {
		ctx = nil
	}
	return
}

func (ctx *Context) Clean() {
	log.Info("Cleaning context")
	ctx.profileDb.Close()
}

func (ctx *Context) HandleRegistration(nf *models.NfProfile) (err error) {
	//a simple implementation
	err = ctx.profileDb.AddNfProfile(nf)
	return
}

func (ctx *Context) HandleHeartbeat(nfid string) (err error) {
	err = ctx.profileDb.UpdateHeartbeat(nfid)
	return
}
