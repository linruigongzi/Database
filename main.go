package main

import (
	"github.reason.com/redis"
	"fmt"
)

func main() {
	redis.BasicGetSet()
	fmt.Println("")
	redis.ExpirationGetSet()
}
