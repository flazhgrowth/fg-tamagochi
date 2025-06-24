package cache

import "time"

type CacheConfig struct {
	Addr        string
	Password    string
	DB          int
	DialTimeout time.Duration
	ReadTimeout time.Duration
}
