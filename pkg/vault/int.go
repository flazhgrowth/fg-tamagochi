package vault

func (vault *Vault) GetIntWithDefault(key string, defaultVal int) int {
	val := vault.app.GetInt(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (vault *Vault) GetInt32WithDefault(key string, defaultVal int32) int32 {
	val := vault.app.GetInt32(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (vault *Vault) GetInt64WithDefault(key string, defaultVal int64) int64 {
	val := vault.app.GetInt64(key)
	if val == 0 {
		return defaultVal
	}

	return val
}
