package config

func (cfg *Config) GetFloat64WithDefault(key string, defaultVal float64) float64 {
	val := cfg.viper.GetFloat64(key)
	if val == 0 {
		return defaultVal
	}

	return val
}
