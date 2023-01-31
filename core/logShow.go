package core

import (
	log "github.com/sirupsen/logrus"
)

func (c *BaseConf) Show(err error, texts ...string) {
	if err != nil {
		log.Panicln(err.Error())
	} else {
		for _, v := range texts {
			log.Println(v)
		}
	}
}
