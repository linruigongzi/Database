package main

import (
	_ "github.com/reason/context"
	rsflag "github.com/reason/flags"
	_ "github.com/reason/rabbitMQ"
	_ "github.com/reason/redis"
)

func main() {
	rsflag.FlagTest()
}
