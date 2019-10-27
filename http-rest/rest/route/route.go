package route

import (
	"github.com/ShiinaOrez/GoDistributedCache/http-rest/rest/handler"
	"net/http"
)

type Router struct {
	Map map[string]http.Handler
}

var CacheRouter Router

func Init() {
	CacheRouter = Router{
		Map: make(map[string]http.Handler),
	}
	CacheRouter.Map["/cache/"] = &handler.CacheHandler
}
