package main

import (
	"fmt"

	"github.reason.com/redis"
)

func main() {
	redis.BasicGetSet()
	fmt.Println("")
	redis.ExpirationGetSet()
	fmt.Println("")
	redis.CheckExist()
	fmt.Println()
	redis.DeleteKey()
}
