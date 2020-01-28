package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ServConfig struct {
	Network map[string]*Network `yaml:"network"`
	Recaptcha map[string]*Recaptcha `yaml:"recaptcha"`
	Security map[string]*Security `yaml:"security"`
	Content map[string]*Content `yaml:"content"`
}

type Network struct {
	listen string `yaml:"listen"`
	mongodb_url string `yaml:"mongodb_url"`
}

type Recaptcha struct {
	enable bool `yaml:"enable"`
	secret_key string `yaml:"secret_key"`
	site_key string	`yaml:"site_key"`
}

type Security struct {
	master_key string `yaml:"master_key"`
	encryption_key string `yaml:"encryption_key"`
	encryption_nonce string `yaml:"encryption_nonce"`
}

type Content struct {
	detect_abuse bool `yaml:"detect_abuse"`
	expire_hrs int `yaml:"expire_hrs"`
}

func checkConfig(servConf ServConfig) (int) {

}

func loadConfig(filePath string) (ServConfig, error) {
	var conf = ServConfig{}
	yamlFd, err := ioutil.ReadFile(filePath)
	log.Printf("Loaded Config: %s \n", filePath)
	if err != nil {
		log.Fatalf("IO Read Error. %#v", err)
		return conf, err
	}
	err = yaml.Unmarshal(yamlFd, &conf)
	if err != nil {
		log.Fatalf("YAML Decode Error. %#v", err)
		return conf, err
	}
	var status int = checkConfig(conf)
	switch status{
	case 0:
	case 1:
		err = errors.New("Security Option not met standard requirement.")
	case 2:
		err = errors.New("MongoDB URI not set or invalid URI schema.")
	case 3:
		err = errors.New("The value is invalid or conflicted in settings.")
	}
	return conf, err
}
