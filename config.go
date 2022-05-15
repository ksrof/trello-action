package config

// ConfigOptions is an alias for a function that will take a pointer to Config and modify it.
type ConfigOptions func(c *Config) error

// Config holds the configuration to interact with both the Github and Trello APIs.
type Config struct {
	methods []string
	paths   []string
	params  map[string]string
}

// WithMethods sets the http request methods.
func WithMethods(method ...string) ConfigOptions {
	return func(c *Config) error {
		c.methods = append(c.methods, method...)
		return nil
	}
}

// WithPaths sets the http request paths.
func WithPaths(path ...string) ConfigOptions {
	return func(c *Config) error {
		c.paths = append(c.paths, path...)
		return nil
	}
}

// WithParams sets the http requests url/query params.
func WithParams(params map[string]string) ConfigOptions {
	return func(c *Config) error {
		c.params = params
		return nil
	}
}

// NewConfig creates a new configuration object with the given options.
func NewConfig(opts ...ConfigOptions) (*Config, error) {
	c := &Config{}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
