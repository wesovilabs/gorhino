package configuration

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/op/go-logging"
	"github.com/wesovilabs/taurus/util"
	"time"
)

var log = logging.MustGetLogger("taurus")
var configFile = "./config.toml"

//Properties - properties structure
type Properties struct {
	Server   Server
	Database Database
}

//Server - server configuration properties
type Server struct {
	Hostname     string
	Port         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

//Database - database configuration propeties
type Database struct {
	Path string
}

//LoadProperties - loading properties from configuration file
func LoadProperties() Properties {
	log.Info("Loading configuration file")
	var conf Properties
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		fmt.Println(err)
		log.Fatal("The program couldn't load the configuration properties.")
	}
	return conf
}

//ServerAddress - it returns the server address
func (server Server) ServerAddress() string {
	return util.ConcatStrings([]string{server.Hostname, ":", server.Port})
}
