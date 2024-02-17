package main

import (
	"crypto/tls"
	"time"

	"github.com/go-redis/redis"
)

type client struct {
	redis *redis.Client
}

type clientoptions func(*client)

// createDefaultClient returns a default redis client
// redis.Options are set with functional arguments
func createDefaultClient(opts ...clientoptions) *client {
	c := &client{redis.NewClient(&redis.Options{})}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func withAddr(addr string) clientoptions {
	return func(c *client) {
		c.redis.Options().Addr = addr
	}
}

func withPassword(pw string) clientoptions {
	return func(c *client) {
		c.redis.Options().Password = pw
	}
}

func withDialTimeout(t time.Duration) clientoptions {
	return func(c *client) {
		c.redis.Options().DialTimeout = t
	}
}

func withReadTimeout(t time.Duration) clientoptions {
	return func(c *client) {
		c.redis.Options().ReadTimeout = t
	}
}

func withTLS(useTLS, skipVerify bool) clientoptions {
	if useTLS {
		return func(c *client) {
			c.redis.Options().TLSConfig = &tls.Config{
				InsecureSkipVerify: skipVerify,
			}
		}
	}
	return func(c *client) {} // omitting TLSConfig disables TLS
}

func (c *client) redisPing() (string, error) {
	return c.redis.Ping().Result()
}
