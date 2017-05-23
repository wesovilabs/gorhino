package connector

import (
	"github.com/wesovilabs/taurus/model"
)

//Driver interface
type Driver interface {
	LoadConfiguration() (services []model.Service, err error)
}
