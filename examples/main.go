package main

import (
	"os"

	"github.com/dinglebit/log"
)

func main() {
	log.Set(log.NewColor(os.Stdout))

	log.WithFields(map[string]interface{}{
		"errored": "for this reason",
	}).Errorf("this is an error message")

	log.WithFields(map[string]interface{}{
		"warned": "for this reason",
	}).Warnf("this is a warn message")

	log.WithFields(map[string]interface{}{
		"infoed": "for this reason",
	}).Infof("this is an info message")

	log.WithFields(map[string]interface{}{
		"debugged": "for this reason",
	}).Debugf("this is a debug message")
}
