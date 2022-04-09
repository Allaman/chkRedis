## chkRedis

A small Go program to verify the connection to a [Redis](https://redis.io/) in-memory data store.

- run `./chkRedis -h` for available options and defaults
- Tested against Redis 6 with TLS and password authentication
- Executes a `PING` command to verify the connection
- Not compatible with Redis RBAC

See my [blog post]() why I wrote chkRedis.
