package checker

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/wesovilabs/taurus/connector"
	"github.com/wesovilabs/taurus/model"
	"time"
)

var log = logging.MustGetLogger("taurus_checker")
var domains []model.Domain

//RegisterDomain registering domain to be listened
func RegisterDomain(domain model.Domain) error {
	log.Info("Registering domain.")
	domains = append(domains, domain)
	schedule(domain, 10*time.Second)
	return nil
}

func unRegisterDomain(domainsUID string) error {
	return nil
}

func schedule(domain model.Domain, delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			var conf model.MongoDBConfiguration = model.MongoDBConfiguration(domain.Settings)
			var driver *connector.MongoDB = connector.NewMongoDB(conf.Hostname, conf.Database, conf.Collection)
			go listenAppConfiguration(driver)
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()
	return stop
}

func listenAppConfiguration(driver connector.Driver) {
	log.Info("Checking configuration for domain....")
	configuration, error := driver.LoadConfiguration()
	if error != nil {

	}

	fmt.Println(configuration)
}
