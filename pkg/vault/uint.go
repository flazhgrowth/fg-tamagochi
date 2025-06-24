package vault

func (vault *Vault) GetUintWithDefault(key string, defaultVal uint) uint {
	val := vault.app.GetUint(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (vault *Vault) GetUint16WithDefault(key string, defaultVal uint16) uint16 {
	val := vault.app.GetUint16(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (vault *Vault) GetUint32WithDefault(key string, defaultVal uint32) uint32 {
	val := vault.app.GetUint32(key)
	if val == 0 {
		return defaultVal
	}

	return val
}

func (vault *Vault) GetUint64WithDefault(key string, defaultVal uint64) uint64 {
	val := vault.app.GetUint64(key)
	if val == 0 {
		return defaultVal
	}

	return val
}
