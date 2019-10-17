package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// BasicGetSet test redigo connect SET GET
func BasicGetSet() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superLee")
	if err != nil {
		fmt.Println("redis set failed: ", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}

// ExpirationGetSet test redigo connect SET GET with expiration
func ExpirationGetSet() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superLee", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed: ", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
