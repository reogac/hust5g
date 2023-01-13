package producer

import (
	"encoding/json"
	"hust5g/sbi/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Producer) OnRegister(c *gin.Context) {
	log.Info("Receive a registration request")
	var nf models.NfProfile
	var err error
	var body []byte
	if body, err = ioutil.ReadAll(c.Request.Body); err == nil {
		if err = json.Unmarshal(body, &nf); err == nil {
			//log pretty NfProfile
			prettynf, _ := json.MarshalIndent(&nf, "    ", "  ")
			log.Infof("%s", string(prettynf))

			if err = p.backend.Context().HandleRegistration(&nf); err == nil {
				c.String(http.StatusAccepted, "Nf %s is registered\n", nf.Id)
			}
		}
	}
	if err != nil {
		//handle error
		log.Infof("Registration failed: error = %s\n", err.Error())
		c.AbortWithError(404, err)
	}
}

func (p *Producer) OnHeartbeat(c *gin.Context) {
	log.Info("Receive a heartbeat request")
	if id := c.Param("nfid"); len(id) > 0 {
		log.Infof("Heartbeat is from %s", id)

		if err := p.backend.Context().HandleHeartbeat(id); err != nil {
			c.AbortWithError(400, err)
		} else {
			c.String(http.StatusAccepted, "the heartbeat from %s is received\n", id)
		}
	} else {
		c.AbortWithStatus(404)
	}
}
