package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

type ServConfig struct {
	Network   Network   `yaml:"network"`
	Recaptcha Recaptcha `yaml:"recaptcha"`
	Security  Security  `yaml:"security"`
	Content   Content   `yaml:"content"`
}

type Network struct {
	Listen      string `yaml:"listen"`
	Mongodb_url string `yaml:"mongodb_url"`
}

type Recaptcha struct {
	Enable     bool   `yaml:"enable"`
	Secret_key string `yaml:"secret_key"`
	Site_key   string `yaml:"site_key"`
}

type Security struct {
	Master_key       string `yaml:"master_key"`
	Encryption_key   string `yaml:"encryption_key"`
	Encryption_nonce string `yaml:"encryption_nonce"`
}

type Content struct {
	Detect_abuse bool `yaml:"detect_abuse"`
	Expire_hrs   int  `yaml:"expire_hrs"`
}

func CheckConfig(servConf ServConfig) int {
	isDBURI := strings.Contains(servConf.Network.Mongodb_url, "mongodb")
	if !isDBURI {
		return 2
	}
	if servConf.Recaptcha.Enable {
		isCAPTsec := len(servConf.Recaptcha.Secret_key) == 40
		isCAPTsit := len(servConf.Recaptcha.Site_key) == 40
		if !(isCAPTsec && isCAPTsit) {
			return 3
		}
	}
	if (servConf.Content.Expire_hrs > 24) || (servConf.Content.Expire_hrs < 0) {
		return 3
	}
	isMasterKeySec := len(servConf.Security.Master_key) >= 12
	isEncrKeySec := len(servConf.Security.Encryption_key) == 32
	isEncrNonceSec := len(servConf.Security.Encryption_nonce) == 12
	if !(isMasterKeySec || isEncrKeySec || isEncrNonceSec) {
		return 1
	}
	return 0
}

func LoadConfig(filePath string) (ServConfig, error) {
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
	switch status {
	case 0:
		log.Println("Config Validation successfully finished!")
	case 1:
		err = errors.New("Security Option not met standard requirement.")
	case 2:
		err = errors.New("MongoDB URI not set or invalid URI schema.")
	case 3:
		err = errors.New("The value is invalid or conflicted in settings.")
	}
	return conf, err
}
