package environment

import "time"

type setting struct {
	Application struct {
		ContextRequest time.Duration `envconfig:"CONTEXT_REQUEST" default:"2.1s"`
	}

	Server struct {
		Context      string        `envconfig:"SERVER_CONTEXT" default:"petshop-bff-mobile"`
		Port         string        `envconfig:"PORT" default:"9997" required:"true" ignored:"false"`
		ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"10s"`
		WriteTimeout time.Duration `envconfig:"READ_TIMEOUT" default:"10s"`
	}

	Redis struct {
		Addr        string        `envconfig:"REDIS_ADDR" default:"localhost:6379"`
		Password    string        `envconfig:"REDIS_PASSWORD"`
		DB          int           `envconfig:"REDIS_DB" default:"0"`
		PoolSize    int           `envconfig:"POOL_SIZE" default:"100"`
		ReadTimeout time.Duration `envconfig:"READ_TIMEOUT" default:"2s"`
	}
}

var Setting setting
