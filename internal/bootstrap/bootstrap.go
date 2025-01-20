package bootstrap

var config = Config{}

func init() {
	setEnvironment()
}

func Execute() Config {
	return config
}
