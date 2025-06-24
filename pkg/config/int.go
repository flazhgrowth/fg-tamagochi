package config

func (cfg Config) GetIntWithDefault(key string, defaultVal int) int {
	val := cfg.viper.GetInt(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (cfg Config) GetInt32WithDefault(key string, defaultVal int32) int32 {
	val := cfg.viper.GetInt32(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (cfg Config) GetInt64WithDefault(key string, defaultVal int64) int64 {
	val := cfg.viper.GetInt64(key)
	if val == 0 {
		return defaultVal
	}

	return val
}
