package main

import (
	"dbsample/redis"
	"fmt"
)

func main() {
	redis.BasicGetSet()
	fmt.Println("\n")
	redis.ExpirationGetSet()
}
