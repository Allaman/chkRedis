package main

import (
	"fmt"
	"time"
)

type CLI struct {
	Version     versionCmd         `cmd:"" help:"Show version information"`
	Check       checkConnectionCmd `cmd:"" default:"1" help:"Check connection to redis (default command)"`
	Addr        string             `default:"localhost:6379" short:"a" help:"Address to connect to in the form host:port"`
	Password    string             `default:"" short:"p" help:"Redis password"`
	DialTimeout time.Duration      `default:"5s" help:"Set DialTimout"`
	ReadTimeout time.Duration      `default:"3s" help:"Set ReadTimeout"`
	TLS         bool               `help:"Connect to a TLS enabled Redis"`
	SkipVerify  bool               `short:"x" help:"Insecure! Accepty every certificate"`
}

type versionCmd struct {
	Version string
}

func (c *versionCmd) Run() error {
	fmt.Println(Version)
	return nil
}

type checkConnectionCmd struct{}

func (c *checkConnectionCmd) Run() error {
	return checkConnection()
}
