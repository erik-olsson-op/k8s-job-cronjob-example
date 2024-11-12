package main

import (
	"github.com/erik-olsson-op/consumer/server"
	"github.com/erik-olsson-op/shared/utils"
	"sync"
)

func main() {
	port := utils.GetEnv("CONSUMER_PORT")
	var wg sync.WaitGroup
	wg.Add(1)
	server.Init(port, &wg)
	wg.Wait()
}
