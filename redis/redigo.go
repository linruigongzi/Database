package rsredis

import (
	"encoding/json"
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

// RedioBasicGetSet test redigo connect SET GET
func RedioBasicGetSet() {
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

// RedioExpirationGetSet test redigo connect SET GET with expiration
func RedioExpirationGetSet() {
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

// RedioAppendExpiration test append expiration later
func RedioAppendExpiration() {
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

	n, _ := c.Do("EXPIRE", "mykey", 2)
	if n == int64(1) {
		fmt.Println("success")
	}

	username, err = redis.String(c.Do("GET", "mykey"))
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

// RedioCheckExist test redigo check Key Exist
func RedioCheckExist() {
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

// RedioDeleteKey test redis delete key operation
func RedioDeleteKey() {
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

// RedioJSONValue test redis GET/SET json value
func RedioJSONValue() {
	c, err := redis.Dial(NetWork, EndPoint)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
		return
	}

	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(valueGet, &imapGet)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}

// RedioListValue test redist List values
func RedioListValue() {
	c, err := redis.Dial(NetWork, EndPoint)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	// multitest will add more values
	// c.Do("DEL", "listkey")

	_, err = c.Do("LPUSH", "listkey", "redis")
	if err != nil {
		fmt.Println("redis set failed: ", err)
	}

	_, err = c.Do("LPUSH", "listkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed: ", err)
	}

	_, err = c.Do("LPUSH", "listkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed: ", err)
	}

	values, _ := redis.Values(c.Do("LRANGE", "listkey", "0", "100"))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
}
