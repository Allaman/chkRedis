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

// createClient returns a redis client
// redis.Options are set with functional arguments
func createClient(opts ...clientoptions) *client {
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
	} else {
		return func(c *client) {} // no TLSConfig disables TLS
	}
}

func (c *client) redisPing() (string, error) {
	pong, err := c.redis.Ping().Result()
	if err != nil {
		return "", err
	} else {
		return pong, nil
	}
}
