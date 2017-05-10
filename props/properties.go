package props

import (
	"gihub.com/wesovilabs/gorhino/util"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/op/go-logging"
	"time"
)

var configFile = "./config.toml"
var log = logging.MustGetLogger("swat_demo_rest_api")

//Properties - properties structure
type Properties struct {
	Server Server
}

//Server - server structure
type Server struct {
	Hostname     string
	Port         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
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
