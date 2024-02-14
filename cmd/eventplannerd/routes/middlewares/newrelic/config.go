package newrelic

import (
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewrelicConfig() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Rest-api"),
		newrelic.ConfigLicense("81bf887466d7ddeacad1db997c4e7952FFFFNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if nil != err {
		return nil, err
	}

	if err = app.WaitForConnection(5 * time.Second); nil != err {
		return nil, err
	}

	return app, nil
}
