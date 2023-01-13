package db

import (
	"fmt"
	"hust5g/sbi/models"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "db"})
}

type Database struct {
	profiles map[string]*models.NfProfile
}

//load database
func Load(dbname string) (db *Database, err error) {
	db = &Database{
		profiles: make(map[string]*models.NfProfile),
	}
	//add code to open a database
	log.Infof("a dummy database is loaded instead of %s", dbname)
	return
}

func (db *Database) AddNfProfile(nf *models.NfProfile) (err error) {
	//just an simplified implementation
	log.Infof("Write NF %s profile to database", nf.Id)
	db.profiles[nf.Id] = nf
	return
}

func (db *Database) UpdateHeartbeat(id string) (err error) {
	if _, ok := db.profiles[id]; ok {
		log.Infof("Update heartbeat for %s", id)
	} else {
		err = fmt.Errorf("No Nf with id=%s is found", id)
		log.Info(err.Error())
	}
	return
}

func (db *Database) Search(query *models.NfQuery) (results []*models.NfProfile) {
	//Do querying here
	return
}

func (db *Database) Close() {
	//TODO: close the database here
	log.Info("Database is closed")
}

//add more methods to read/write the database
