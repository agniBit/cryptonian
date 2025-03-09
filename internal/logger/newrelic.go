package logger

import (
	"sync"

	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var initNewRelicOnce = &sync.Once{}
var newRelicApp *newrelic.Application

func InitNewRelic(cfg *cfg.Config) {
	initNewRelicOnce.Do(func() {
		newRelic, err := newrelic.NewApplication(
			newrelic.ConfigAppName(cfg.NewRelic.AppName),
			newrelic.ConfigLicense(cfg.NewRelic.LicenseKey),
			newrelic.ConfigDistributedTracerEnabled(true),
		)

		if err != nil {
			if cfg.Server.Environment == "prod" {
				panic(err)
			}
		}

		newRelicApp = newRelic
	})
}

func GetApp() *newrelic.Application {
	return newRelicApp
}
