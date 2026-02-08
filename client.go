package main

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/redis/go-redis/v9"
)

type client struct {
	redis *redis.Client
}

type clientoptions func(*redis.Options)

// createDefaultClient returns a default redis client
// redis.Options are set with functional arguments
func createDefaultClient(opts ...clientoptions) *client {
	options := &redis.Options{}
	for _, opt := range opts {
		opt(options)
	}
	return &client{redis: redis.NewClient(options)}
}

func withAddr(addr string) clientoptions {
	return func(o *redis.Options) {
		o.Addr = addr
	}
}

func withUsername(username string) clientoptions {
	return func(o *redis.Options) {
		o.Username = username
	}
}

func withPassword(pw string) clientoptions {
	return func(o *redis.Options) {
		o.Password = pw
	}
}

func withDialTimeout(t time.Duration) clientoptions {
	return func(o *redis.Options) {
		o.DialTimeout = t
	}
}

func withReadTimeout(t time.Duration) clientoptions {
	return func(o *redis.Options) {
		o.ReadTimeout = t
	}
}

func withTLS(useTLS, skipVerify bool) clientoptions {
	if useTLS {
		return func(o *redis.Options) {
			o.TLSConfig = &tls.Config{
				InsecureSkipVerify: skipVerify,
			}
		}
	}
	return func(o *redis.Options) {} // omitting TLSConfig disables TLS
}

func (c *client) redisPing() (string, error) {
	return c.redis.Ping(context.Background()).Result()
}
