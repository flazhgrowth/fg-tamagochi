package featureflag

import (
	"github.com/spf13/viper"
)

type FeatureFlag struct {
	viper *viper.Viper
}

var (
	ff *FeatureFlag
)

func New() error {
	v := viper.New()
	v.SetConfigFile("./etc/featureflag/featureflag.yaml")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	ff = &FeatureFlag{
		viper: v,
	}

	return nil
}

func GetFeatureFlag() *FeatureFlag {
	return ff
}

func (ff *FeatureFlag) GetViper() *viper.Viper {
	return ff.viper
}

func (ff *FeatureFlag) IsEnabled(key string) bool {
	return ff.viper.GetBool(key)
}
