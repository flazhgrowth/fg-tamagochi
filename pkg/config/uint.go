package config

func (cfg Config) GetUintWithDefault(key string, defaultVal uint) uint {
	val := cfg.viper.GetUint(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (cfg Config) GetUint16WithDefault(key string, defaultVal uint16) uint16 {
	val := cfg.viper.GetUint16(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (cfg Config) GetUint32WithDefault(key string, defaultVal uint32) uint32 {
	val := cfg.viper.GetUint32(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (cfg Config) GetUint64WithDefault(key string, defaultVal uint64) uint64 {
	val := cfg.viper.GetUint64(key)
	if val == 0 {
		return defaultVal
	}

	return val
}
