package util

import (
	"bytes"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("swat_demo_rest_api/util")

//ConcatStrings - Concat an array of strings
func ConcatStrings(args []string) string {
	log.Debug("Concatenating strings arrays")
	var str bytes.Buffer
	for _, l := range args {
		str.WriteString(l)
	}
	return str.String()
}
