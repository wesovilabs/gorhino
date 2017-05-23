package configuration

import (
	"github.com/op/go-logging"
	"os"
)

//InitializeLogger setup logging configuration
func InitializeLogger() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}
