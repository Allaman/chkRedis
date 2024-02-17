package main

import (
	"log"
	"log/slog"

	"github.com/alecthomas/kong"
)

// will be overwritten in release pipeline
var Version = "dev"

var c *client

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("chkRedis"),
		kong.Description("Test the connection to Redis"))
	c = createDefaultClient(
		withAddr(cli.Addr),
		withPassword(cli.Password),
		withDialTimeout(cli.DialTimeout),
		withReadTimeout(cli.ReadTimeout),
		withTLS(cli.TLS, cli.SkipVerify),
	)
	err := ctx.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func checkConnection() error {
	res, err := c.redisPing()
	if err != nil {
		slog.Error("error pinging Redis: ", err)
		return err
	}
	slog.Info(res)
	return nil
}
