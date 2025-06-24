package config

func (cfg Config) IsEnvProduction() bool {
	return cfg.Env() == "production"
}

func (cfg Config) Env() string {
	return cfg.GetStringWithDefault("env", "local")
}
