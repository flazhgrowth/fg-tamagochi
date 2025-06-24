package vault

func (vault *Vault) GetFloat64WithDefault(key string, defaultVal float64) float64 {
	val := vault.app.GetFloat64(key)
	if val == 0 {
		return defaultVal
	}

	return val
}
