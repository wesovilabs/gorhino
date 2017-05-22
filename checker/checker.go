package checker

import (
	"time"
	"github.com/wesovilabs/taurus/connector"
)

var domains []string

var apps = make(map[string]connector.Driver)

func registerDomain(domainUID string) error {
	domains = append(domains,domainUID)
	return nil
}

func unRegisterDomain(domainsUID string) error {
	return nil
}

func schedule(appUID string, driver connector.Driver, delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			listenAppConfiguration(appUID, driver)
			select {
			case <- time.After(delay):
			case <- stop:
				return
			}
		}
	}()
	return stop
}

func listenAppConfiguration(appUID string, driver connector.Driver){
	go updateAppConfiguration(appUID,driver.Configuration())
}

func updateAppConfiguration(appUID string, appConfiguration string){

}