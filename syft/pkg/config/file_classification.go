package config

import (
	"github.com/anchore/syft/syft/source"
	"github.com/spf13/viper"
)

type FileClassification struct {
	Cataloger CatalogerOptions `yaml:"cataloger" json:"cataloger" mapstructure:"cataloger"`
}

func (cfg FileClassification) loadDefaultValues(v *viper.Viper) {
	v.SetDefault("file-classification.cataloger.enabled", catalogerEnabledDefault)
	v.SetDefault("file-classification.cataloger.scope", source.SquashedScope)
}

func (cfg *FileClassification) parseConfigValues() error {
	return cfg.Cataloger.parseConfigValues()
}
