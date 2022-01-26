package gocelery

// Option sets up a Config.
type Option func(*Config)

// WithIgnoreResult disables a task result storage.
// By default results are stored in the backend.
func WithIgnoreResult() Option {
	return func(c *Config) {
		c.ignoreResult = true
	}
}

// Config represents Celery client configuration.
type Config struct {
	// ignoreResult if set to false (default) the result of a task will be stored in the backend.
	// If set to true the result won't be stored.
	ignoreResult bool
}
