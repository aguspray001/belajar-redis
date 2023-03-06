package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.0.2.15:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0,
	})
	var a int
	a = 1
	fmt.Print(a)

	err := rdb.Set(ctx, "testing-value", 99, 5*time.Second).Err()

	if err != nil {
		panic(err)
	}
	time.Sleep(4 * time.Second)
	val, err := rdb.Get(ctx, "testing-value").Result()

	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
