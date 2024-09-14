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
	if os.Getenv("LOCAL_E2E_TEST") != "" {
		return os.Getenv("LOCAL_E2E_TEST") == "true"
	}
	switch os.Getenv("APP_STAGE_TYPE") {
	case "dev", "localdev":
		return true
	default:
		return false
	}
})
