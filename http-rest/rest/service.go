package rest

import (
	"fmt"
	"net/http"

	"github.com/ShiinaOrez/GoDistributedCache/http-rest/cache"
	"github.com/ShiinaOrez/GoDistributedCache/http-rest/constvar"
	"github.com/ShiinaOrez/GoDistributedCache/http-rest/rest/handler"
	"github.com/ShiinaOrez/GoDistributedCache/http-rest/rest/route"
)

type RESTServer struct {
	CacheService *cache.CacheService
}

var server *RESTServer

func (server *RESTServer) start() {
	fmt.Println("[start]")
	route.Init()
	handler.SetCacheService(server.CacheService)
	for path, handler := range route.CacheRouter.Map {
		fmt.Println("[handler]:", path, handler)
		http.Handle(path, handler)
	}
	fmt.Println("Listen on: localhost" + constvar.Port + "...")
	http.ListenAndServe(constvar.Port, nil)
}

func StartWithCacheService(cacheService *cache.CacheService) {
	server = &RESTServer{
		CacheService: cacheService,
	}
	server.start()
}
