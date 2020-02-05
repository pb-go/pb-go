package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	CurrentVer string = "v0.1.0"
	CurrentDSN string = "https://72dd7f93900d4742a436a525692a13ed@sentry.io/2124482"
)

var (
	ServConf ServConfig
)

type ServConfig struct {
	// server side config file
	Network   Network   `yaml:"network"`
	Recaptcha Recaptcha `yaml:"recaptcha"`
	Security  Security  `yaml:"security"`
	Content   Content   `yaml:"content"`
}

type Network struct {
	// subconfig about network
	Listen     string `yaml:"listen"`
	Host       string `yaml:"host"`
	MongodbUrl string `yaml:"mongodb_url"`
}

type Recaptcha struct {
	// subconfig about recaptcha
	Enable    bool   `yaml:"enable"`
	SecretKey string `yaml:"secret_key,omitempty"`
	SiteKey   string `yaml:"site_key,omitempty"`
}

type Security struct {
	// subconfig about data encryption and administration
	MasterKey       string `yaml:"master_key"`
	EncryptionKey   string `yaml:"encryption_key"`
	EncryptionNonce string `yaml:"encryption_nonce"`
}

type Content struct {
	// subconfig about content abusing detection
	DetectAbuse       bool `yaml:"detect_abuse"`
	ExpireHrs         int  `yaml:"expire_hrs"`
	AllowBase64encode bool `yaml:"allow_b64enc"`
}

func CheckConfig(servConf ServConfig) int {
	// detect uri validity
	isDBURI := strings.Contains(servConf.Network.MongodbUrl, "mongodb")
	if !isDBURI {
		return 2
	}
	// check needed config for recaptcha
	if servConf.Recaptcha.Enable {
		isCAPTsec := len(servConf.Recaptcha.SecretKey) == 40
		isCAPTsit := len(servConf.Recaptcha.SiteKey) == 40
		if !(isCAPTsec && isCAPTsit) {
			return 3
		}
	}
	// check max expire
	if (servConf.Content.ExpireHrs > 24) || (servConf.Content.ExpireHrs < 0) {
		return 3
	}
	// check cryptography requirement
	isMasterKeySec := len(servConf.Security.MasterKey) >= 12
	isEncrKeySec := len(servConf.Security.EncryptionKey) == 32
	isEncrNonceSec := len(servConf.Security.EncryptionNonce) == 12
	if !(isMasterKeySec || isEncrKeySec || isEncrNonceSec) {
		return 1
	}
	return 0
}

func FileExist(filepath string) bool {
	// check if file exists, utils
	info, err := os.Stat(filepath)
	return err == nil && !info.IsDir()
}

func LoadConfig(filePath string) (ServConfig, error) {
	// config load function
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
	var status int = CheckConfig(conf)
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
