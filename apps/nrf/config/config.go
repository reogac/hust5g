package config

import (
	"encoding/json"
	"hust5g/sbi"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "config"})
}

type Config struct {
	Sbi    sbi.ServerConfig `json:"sbi"`
	DbName string           `json:"dbname"` //database endpoint
}

func Load(f string) (conf Config, err error) {
	log.Infof("read configuration: %s", f)
	var buf []byte
	if buf, err = ioutil.ReadFile(f); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &conf)
	}
	return
}
