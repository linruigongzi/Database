package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	// NetWork specify the redis network
	NetWork = "tcp"
	// EndPoint specify the redis endpoint
	EndPoint = "127.0.0.1:6379"
)

// BasicGetSet test redigo connect SET GET
func BasicGetSet() {
	c, err := redis.Dial(NetWork, EndPoint)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superLee")
	if err != nil {
		fmt.Println("redis set failed: ", err)
		return
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
	c, err := redis.Dial(NetWork, EndPoint)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superLee", "EX", "1")
	if err != nil {
		fmt.Println("redis set failed: ", err)
		return
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(3 * time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}

// CheckExist test redigo check Key Exist
func CheckExist() {
	c, err := redis.Dial(NetWork, EndPoint)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superLee")
	if err != nil {
		fmt.Println("redis set failed: ", err)
		return
	}

	isKeyExist, err := redis.Bool(c.Do("EXISTS", "nokey"))
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("exists or not: %v \n", isKeyExist)
	}
}

// DeleteKey test redis delete key operation
func DeleteKey() {
	c, err := redis.Dial(NetWork, EndPoint)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superLee")
	if err != nil {
		fmt.Println("redis set failed: ", err)
		return
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	_, err = c.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delete failed: ", err)
		return
	}

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

}
