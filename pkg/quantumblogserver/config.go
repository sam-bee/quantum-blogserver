package quantumblogserver

import (
	"log"
	"os"
	"strconv"
)

type ApplicationConfig struct {
	certs struct {
		certpem string
		keypem  string
	}
	serverDetails struct {
		portNo int
	}
	staticSiteDir string
	logFile       string
}

func BuildConfig(settings map[string]interface{}) ApplicationConfig {
	c := ApplicationConfig{}
	c.certs.certpem = getSetting("path_to_cert", settings)
	c.certs.keypem = getSetting("path_to_key", settings)
	c.serverDetails.portNo = getIntSetting("port_no", settings)
	c.staticSiteDir = getSetting("static_site_dir", settings)
	c.logFile = getSetting("log_file", settings)
	return c
}

func (c *ApplicationConfig) getCertPem() string {
	return c.certs.certpem
}

func (c *ApplicationConfig) getKeyPem() string {
	return c.certs.keypem
}

func (c *ApplicationConfig) getPortNo() int {
	return c.serverDetails.portNo
}

func (c *ApplicationConfig) getStaticSiteDir() string {
	return c.staticSiteDir
}

func (c *ApplicationConfig) GetLogFile() string {
	return c.logFile
}

func getSetting(key string, settings map[string]interface{}) string {

	val, ok := settings[key].(string)
	if !ok {
		log.Fatalf("Setting %s not found", key)
		os.Exit(1)
	}

	return val
}

func getIntSetting(key string, settings map[string]interface{}) int {
	s := getSetting(key, settings)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Setting %s is not an integer", key)
		os.Exit(1)
	}
	return i
}
