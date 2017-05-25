package connector

import (
	"github.com/swatlabs/taurus/model"
)

//Driver interface
type Driver interface {
	LoadConfiguration() (services []model.Service, err error)
}
