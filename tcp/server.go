package main

import (
	"github.com/ShiinaOrez/GoDistributedCache/tcp/cache"
	"github.com/ShiinaOrez/GoDistributedCache/tcp/network"
)

func main() {
	cacheService := cache.NewService()
	network.StartWithService(cacheService)
}
