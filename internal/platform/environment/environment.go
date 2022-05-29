package environment

import (
	opportunitiesConfig "EurekaHome/internal/platform/opportunities/config"
	"EurekaHome/internal/platform/opportunities"
	"EurekaHome/internal/platform/sheets"
)

type ScopeConfig struct {
	Port                          string                               `yaml:"Port"`
	TypesSupported                []opportunities.TypesSupported          `yaml:"FinancialOpportunitiesTypesSupported"`
	ProcessorOpportunitiesClient        opportunitiesConfig.ProcessorOpportunitiesClient `yaml:"ProcessorOpportunitiesClient"`
	FinancialOpportunitiesSheetsClient  sheets.Client                           `yaml:"FinancialOpportunitiesSheetsClient"`
}
