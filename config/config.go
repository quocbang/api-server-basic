package config

type Options struct {
	DevMode bool   `long:"dev-mode"`
	Host    string `long:"host" env:"HOST"`
	Port    int    `long:"port" env:"PORT"`
	Config  string `long:"config" env:"CONFIG"`
}

type TLSOptions struct {
	Cert string `long:"tls-cert"`
	Key  string `long:"tls-key"`
}

func (t *TLSOptions) UseTLS() bool {
	return t.Cert != "" && t.Key != ""
}

type Configs struct {
	Postgres    PostgresConfig      `yaml:"postgres"`
	Redis       RedisConfig         `yaml:"redis"`
	Roles       map[string][]string `yaml:"roles"`
	DemoAccount []string            `yaml:"demo-account"`
}

type PostgresConfig struct {
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Schema   string `yaml:"schema"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}
