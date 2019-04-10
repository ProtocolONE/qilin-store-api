package conf

type ServerConfig struct {
	Port             int      `envconfig:"PORT" required:"false" default:"8080"`
	AllowOrigins     []string `envconfig:"ALLOW_ORIGINS" required:"false" default:"*"`
	AllowCredentials bool     `envconfig:"ALLOW_CREDENTIALS" required:"false" default:"false"`
	Debug            bool     `envconfig:"DEBUG" required:"false" default:"false"`
}

type DbConfig struct {
	Connection string `envconfig:"CONNECTION" required:"true" default:"localhost:27017"`
	Name       string `envconfig:"DB_NAME" required:"true" default:"DEVELOP"`
}
