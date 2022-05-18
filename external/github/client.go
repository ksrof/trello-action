package github

type options func(c *config) error

type config struct {
	token string
	user  string
	repo  string
	event string
	id    string
}

func New(opts ...options) (*config, error) {
	c := &config{}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithToken(token string) options {
	return func(c *config) error {
		c.token = token
		return nil
	}
}

func WithUser(user string) options {
	return func(c *config) error {
		c.user = user
		return nil
	}
}

func WithRepo(repo string) options {
	return func(c *config) error {
		c.repo = repo
		return nil
	}
}

func WithEvent(event string) options {
	return func(c *config) error {
		c.event = event
		return nil
	}
}

func WithID(id string) options {
	return func(c *config) error {
		c.id = id
		return nil
	}
}
