package main

import (
	"flag"
	"grpc-go-study/cmd"
	"grpc-go-study/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)
	cmd.NewApp(cfg)
}
