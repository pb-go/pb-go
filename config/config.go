package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// software version, sentry bug tracking id
const (
	CurrentVer string = "v1.0.0"
	CurrentDSN string = "https://72dd7f93900d4742a436a525692a13ed@sentry.io/2124482"
)

// global var for global server config read
var (
	ServConf ServConfig
)

// ServConfig consists of 4 parts
type ServConfig struct {
	Network   Network   `yaml:"network"`
	Recaptcha Recaptcha `yaml:"recaptcha"`
	Security  Security  `yaml:"security"`
	Content   Content   `yaml:"content"`
}

//Network : subconfig of ServConf
type Network struct {
	Listen     string `yaml:"listen"`
	Host       string `yaml:"host"`
	MongodbURL string `yaml:"mongodb_url"`
}

// Recaptcha : subconfig of ServConf
type Recaptcha struct {
	Enable    bool   `yaml:"enable"`
	SecretKey string `yaml:"secret_key,omitempty"`
	SiteKey   string `yaml:"site_key,omitempty"`
}

// Security : subconfig about data encryption and administration
// must fulfill chacha20 standard
type Security struct {
	MasterKey       string `yaml:"master_key"`
	EncryptionKey   string `yaml:"encryption_key"`
	EncryptionNonce string `yaml:"encryption_nonce"`
}

// Content : subconfig about content abusing detection
type Content struct {
	DetectAbuse bool `yaml:"detect_abuse"`
	ExpireHrs   int  `yaml:"expire_hrs"`
}

// CheckConfig : detect uri validity, check needed config for recaptcha, check max expire, check cryptography requirement
func CheckConfig(servConf ServConfig) int {
	isDBURI := strings.Contains(servConf.Network.MongodbURL, "mongodb")
	if !isDBURI {
		return 2
	}
	if servConf.Recaptcha.Enable {
		isCAPTsec := len(servConf.Recaptcha.SecretKey) == 40
		isCAPTsit := len(servConf.Recaptcha.SiteKey) == 40
		if !(isCAPTsec && isCAPTsit) {
			return 3
		}
	}
	if (servConf.Content.ExpireHrs > 24) || (servConf.Content.ExpireHrs < 0) {
		return 3
	}
	isMasterKeySec := len(servConf.Security.MasterKey) >= 12
	isEncrKeySec := len(servConf.Security.EncryptionKey) == 32
	isEncrNonceSec := len(servConf.Security.EncryptionNonce) == 12
	if !(isMasterKeySec || isEncrKeySec || isEncrNonceSec) {
		return 1
	}
	return 0
}

// FileExist : check if file exists, utils
func FileExist(filepath string) bool {
	info, err := os.Stat(filepath)
	return err == nil && !info.IsDir()
}

// LoadConfig : config load function, read from file.
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
	// config validity check.
	var status = CheckConfig(conf)
	switch status {
	case 0:
		log.Println("Config validation successfully finished!")
	case 1:
		err = errors.New("security option not met standard requirement")
	case 2:
		err = errors.New("mongoDB URI not set or invalid URI schema")
	case 3:
		err = errors.New("the value is invalid or conflicted in settings")
	}
	return conf, err
}
