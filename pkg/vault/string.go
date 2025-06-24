package vault

func (vault *Vault) GetStringWithDefault(key string, defaultVal string) string {
	val := vault.app.GetString(key)
	if val == "" {
		return defaultVal
	}

	return val
}
