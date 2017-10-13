package config

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	Server *Server
}

func Load(path string) (*Configuration, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.WithFields(log.Fields{
			"path": path,
		}).Error("Failed reading configuration file")

		return nil, err
	}

	config := &Configuration{}

	err = yaml.Unmarshal(data, config)

	if err != nil {
		log.WithFields(log.Fields{
			"path": path,
		}).Error("Failed unmarshalling configuration")

		return nil, err
	}

	return config, nil
}
