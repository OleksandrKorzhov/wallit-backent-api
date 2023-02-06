package plaid

import (
	plaidSdk "github.com/plaid/plaid-go/v10/plaid"
	"github.com/spf13/viper"
	"log"
	"wallit/internal/config"
)

func ClientId() string {
	return config.GetCriticalConfigParam[string](func() interface{} {
		return viper.Get("PLAID_CLIENT_ID")
	})
}

func Secret() string {
	return config.GetCriticalConfigParam[string](func() interface{} {
		return viper.Get("PLAID_SECRET")
	})
}

const (
	SandboxEnv     = "sandbox"
	DevelopmentEnv = "development"
	ProductionEnv  = "production"
)

func Environment() plaidSdk.Environment {
	value, ok := viper.Get("PLAID_ENVIRONMENT").(string)

	if !ok {
		return plaidSdk.Sandbox
	}

	allowedEnvironments := map[string]plaidSdk.Environment{
		SandboxEnv:     plaidSdk.Sandbox,
		DevelopmentEnv: plaidSdk.Development,
		ProductionEnv:  plaidSdk.Production,
	}

	if allowedEnvironments[value] == "" {
		log.Fatalf("The passed environment is not in the list of recognisable ones - %v", value)
	}

	return allowedEnvironments[value]
}
