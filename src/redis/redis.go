package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"fmt"
)

var (
	c redis.Conn
	err error
	reply interface{}
)

func init() {
	c, err = redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer c.Close()
	c.Do("FLUSHALL")
	stringExample()
}

func stringExample() {
	c.Do("SET", "hello", "world")
	s, err := redis.String(c.Do("GET", "hello"))
	fmt.Printf("%#v %v\n", s, err)
}