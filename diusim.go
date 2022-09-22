package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/mediocregopher/radix/v4"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var iteration int

	for iteration = 0; iteration < 10; iteration++ {
		time.Sleep(1000 * time.Millisecond)
		go startDiu(iteration)
	}

}

func startDiu(iteration int) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	client, err := (radix.PoolConfig{}).New(ctx, "tcp", "127.0.0.1:6379") // or any other client
	for i := 1; i < 100; i++ {
		randomValue := rand.Intn(100)
		if err := client.Do(ctx, radix.Cmd(nil, "SET", "foo"+strconv.Itoa(iteration)+strconv.Itoa(i), strconv.Itoa(randomValue))); err != nil {
			panic(err)
		}
		fmt.Println("foo" + strconv.Itoa(iteration) + strconv.Itoa(i))
	}
	if err != nil {
		panic(err)
	}
}
