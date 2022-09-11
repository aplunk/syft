package config

import (
	"github.com/anchore/syft/syft/pkg/cataloger"
	"github.com/spf13/viper"
)

type Pkg struct {
	Cataloger               CatalogerOptions `yaml:"cataloger" json:"cataloger" mapstructure:"cataloger"`
	SearchUnindexedArchives bool             `yaml:"search-unindexed-archives" json:"search-unindexed-archives" mapstructure:"search-unindexed-archives"`
	SearchIndexedArchives   bool             `yaml:"search-indexed-archives" json:"search-indexed-archives" mapstructure:"search-indexed-archives"`
}

func (cfg Pkg) loadDefaultValues(v *viper.Viper) {
	cfg.Cataloger.loadDefaultValues(v)
	c := cataloger.DefaultSearchConfig()
	v.SetDefault("package.search-unindexed-archives", c.IncludeUnindexedArchives)
	v.SetDefault("package.search-indexed-archives", c.IncludeIndexedArchives)
}

func (cfg *Pkg) parseConfigValues() error {
	return cfg.Cataloger.parseConfigValues()
}
