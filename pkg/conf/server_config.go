package conf

type ServerConfig struct {
	Port             int      `envconfig:"PORT" required:"false" default:"8080"`
	AllowOrigins     []string `envconfig:"ALLOW_ORIGINS" required:"false" default:"*"`
	AllowCredentials bool     `envconfig:"ALLOW_CREDENTIALS" required:"false" default:"false"`
	Debug            bool     `envconfig:"DEBUG" required:"false" default:"false"`
}

type Auth1 struct {
	Issuer       string `envconfig:"ISSUER" required:"true" default:"https://dev-auth1.tst.protocol.one"`
	ClientId     string `envconfig:"CLIENTID" required:"true"`
	ClientSecret string `envconfig:"CLIENTSECRET" required:"true"`
}

type DbConfig struct {
	Host           string `envconfig:"HOST" required:"false" default:"127.0.0.1"`
	Name           string `envconfig:"NAME" required:"false" default:"qilinstoreapi"`
	User           string `envconfig:"USER" required:"false"`
	Password       string `envconfig:"PASSWORD" required:"false"`
	MaxConnections int    `envconfig:"MAX_CONNECTIONS" required:"false" default:"100"`
}

type EventBusConfig struct {
	Connection string `envconfig:"CONNECTION" required:"true" default:"amqp://127.0.0.1:5672"`
}
