package helpers

import (
	"os"
	"strings"

	goosecond "github.com/akm/goosecond"
)

var DemoData = goosecond.NewCondition(func() bool {
	appEnv := os.Getenv("APP_ENV")
	if strings.Contains(appEnv, "test") {
		return false
	}
	if os.Getenv("DEMO_DATA") != "" {
		return os.Getenv("DEMO_DATA") == "true"
	}
	return os.Getenv("APP_STAGE_TYPE") == "local"
})
