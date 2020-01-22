package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	Pool *redis.Pool
)

func init() {
	redisHost := ":6379"
	Pool = newPool(redisHost)
	listenForClose()
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func listenForClose() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		_ = Pool.Close()
		os.Exit(0)
	}()
}

func Get(key string) ([]byte, error) {

	conn := Pool.Get()
	defer conn.Close()
	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func main() {
	test, err := Get("test")
	fmt.Println(test, err)
	ExampleNewClient()
	// Output
	// [] error get key test: redigo: nil returned
	// PONG <nil>
	// 4396 <nil>
	// key: overnote2 not exist
	// 0 redigo: nil returned
}

func ExampleNewClient() {
	c, _ := redis.Dial("tcp", "localhost:6379",
		redis.DialDatabase(0),
	)
	pong, err := redis.String(c.Do("ping"))
	fmt.Println(pong, err)
	// Output: PONG <nil>
	key := "overnote"
	val := 4396
	_, err = c.Do("SET", key, val)
	if err != nil {
		panic(err)
	}
	res, err := redis.Int(c.Do("GET", key))
	if err == redis.ErrNil {
		fmt.Printf("key: %s not exist\n", key)
	} else {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(res, err)
	// 4396 nil
	notExistKey := "overnote2"
	res, err = redis.Int(c.Do("GET", notExistKey))
	if err == redis.ErrNil {
		fmt.Printf("key: %s not exist\n", notExistKey)
		// key: overnote2 not exist
	} else {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(res, err)
	// 0 redigo: nil returned
}
