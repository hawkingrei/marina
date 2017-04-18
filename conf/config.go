package conf

type Config struct {
	HttpAddr string
	Verbose  bool
	Quiet    bool
}

// NewConfig creates a new config
func NewConfig() *Config {
	return &Config{
		//DockerHost:  dockerHost,
		HttpAddr: ":80",
		Verbose:  false,
		Quiet:    false,
	}
}
