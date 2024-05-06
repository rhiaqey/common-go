package env

type Env struct {
	Id        string `env:"ID,notEmpty"`
	Name      string `env:"NAME,notEmpty"`
	Namespace string `env:"NAMESPACE,notEmpty"`
	// PrivateKey  string `env:"PRIVATE_KEY,notEmpty,unset"`
	// PublicKey   string `env:"PUBLIC_KEY,notEmpty,unset"`
	PrivatePort uint16 `env:"PRIVATE_PORT" envDefault:"3001"`
	PublicPort  uint16 `env:"PUBLIC_PORT" envDefault:"3000"`
	// Redis
	RedisMode              string    `env:"REDIS_MODE" envDefault:"standalone"`
	RedisAddress           *string   `env:"REDIS_ADDRESS" envDefault:"localhost:6379"`
	RedisPassword          *string   `env:"REDIS_PASSWORD,unset"`
	RedisDb                *int      `env:"REDIS_DB" envDefault:"0"`
	RedisSentinelMaster    *string   `env:"REDIS_SENTINEL_MASTER"`
	RedisSentinelAddresses *[]string `env:"REDIS_SENTINEL_ADDRESSES" envSeparator:","`
}

func (env *Env) IsStandaloneMode() bool {
	return env.RedisMode == "standalone"
}

func (env *Env) IsSentinelMode() bool {
	if env.RedisMode == "sentinel" {
		if env.RedisSentinelAddresses != nil {
			return len(*env.RedisSentinelAddresses) > 0
		}
	}

	return false
}
