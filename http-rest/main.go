package main

import (
	"github.com/ShiinaOrez/GoDistributedCache/http-rest/cache"
	"github.com/ShiinaOrez/GoDistributedCache/http-rest/rest"
)

func main() {
	cacheService := cache.NewCacheService()
	rest.StartWithCacheService(cacheService)
}
