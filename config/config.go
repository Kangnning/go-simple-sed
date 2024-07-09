package config

type Action uint8

const (
	InsertBefore Action = iota
	InsertAfter
	Delete
	Replace
)

type Config struct {
	FileName  string
	Act       Action
	Pattern   string
	DesString string
}

type Option func(*Config)

func WithFileName(fileName string) Option {
	return func(c *Config) {
		c.FileName = fileName
	}
}

func WithAction(act Action) Option {
	return func(c *Config) {
		c.Act = act
	}
}

func WithPattern(pattern string) Option {
	return func(c *Config) {
		c.Pattern = pattern
	}
}

func WithDesString(des string) Option {
	return func(c *Config) {
		c.DesString = des
	}
}

func New(opt ...Option) *Config {
	conf := &Config{}
	for _, fn := range opt {
		fn(conf)
	}
	return conf
}

func (c *Config) Modify(opt ...Option) {
	for _, fn := range opt {
		fn(c)
	}
}
