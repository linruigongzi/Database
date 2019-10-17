package main

import (
	"database/redis"
	"fmt"
)

func main() {
	redis.BasicGetSet()
	fmt.Println("\n")
	redis.ExpirationGetSet()
}
