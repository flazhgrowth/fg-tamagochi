package config

func (cfg Config) GetStringWithDefault(key string, defaultVal string) string {
	val := cfg.viper.GetString(key)
	if val == "" {
		return defaultVal
	}

	return val
}
