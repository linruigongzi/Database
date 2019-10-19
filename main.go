package main

import (
	// rsredis "github.com/reason/redis"

	_ "github.com/reason/context"
	rsrabbit "github.com/reason/rabbitMQ"
)

func main() {

	rsrabbit.Send("ReasonMQ")

	// rsredis.RedioBasicGetSet()
	// fmt.Println()
	// rsredis.RedioExpirationGetSet()
	// fmt.Println()
	// rsredis.RedioCheckExist()
	// fmt.Println()
	// rsredis.RedioDeleteKey()
	// fmt.Println()
	// rsredis.RedioJSONValue()
	// fmt.Println()
	// rsredis.RedioAppendExpiration()
	// fmt.Println()
	// rsredis.RedioListValue()

	// rsredis.RedisBasicUsage()

	// rscontext.CTWithCancel()
	// fmt.Println()
	// rscontext.CTWithDeadline()
	// fmt.Println()
	// rscontext.CTWithTimeout()
	// fmt.Println()
	// rscontext.CTWithValue()
}
