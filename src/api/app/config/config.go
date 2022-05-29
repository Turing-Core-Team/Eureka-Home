//nolint
package config

import (
	"EurekaHome/internal/platform/constant"
	"EurekaHome/internal/platform/environment"
	"EurekaHome/internal/platform/sheets"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

type Config struct {
	scope     constant.Scope
	envConfig environment.ScopeConfig
}

var configuration *Config

func init() {
	setupEnvironment()
}

func setupEnvironment() {
	fmt.Println("Global Environment Setup")
	scope := SetupScope(os.Getenv("SCOPE"))
	envConfig := setupEnvironmentConfig([]byte(findConfigStringByScope(scope)))

	configuration = &Config{
		scope:     scope,
		envConfig: envConfig,
	}
}

func GetConfig() *Config {
	return configuration
}

func SetupScope(instanceScope string) constant.Scope {
	lowerInstanceScope := strings.ToLower(instanceScope)
	if strings.Contains(lowerInstanceScope, strings.ToLower(constant.ProdScope.Value())) {
		return constant.ProdScope
	} else if strings.Contains(lowerInstanceScope, strings.ToLower(constant.TestScope.Value())) {
		return constant.TestScope
	}else {
		return constant.DevScope
	}
}

func setupEnvironmentConfig(scopeEnvConfigBytes []byte) environment.ScopeConfig {
	cfg := environment.ScopeConfig{}
	err := yaml.Unmarshal(scopeEnvConfigBytes, &cfg)
	if err != nil {
		log.Panicf("error with Scope configs: %v", err)
	}
	return cfg
}

func findConfigStringByScope(scope constant.Scope) string {
	configsByScopeMap := map[constant.Scope]string{
		constant.ProdScope: environment.Production,
		constant.TestScope: environment.Test,
		constant.DevScope:  environment.Development,
	}
	return configsByScopeMap[scope]
}

func (c *Config) Scope() constant.Scope {
	return c.scope
}

func (c *Config) Port() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":" + c.envConfig.Port
}

func (c *Config) SheetsConfig() sheets.Client {
	return c.envConfig.SheetsConfig
}
