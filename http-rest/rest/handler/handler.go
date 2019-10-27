package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ShiinaOrez/GoDistributedCache/http-rest/cache"
)

type cacheHandler struct {
	Service *cache.CacheService
}

var CacheHandler cacheHandler

func Init() {
	CacheHandler = cacheHandler{
		Service: nil,
	}
}

func SetCacheService(service *cache.CacheService) {
	CacheHandler.Service = service
}

func (handler *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// log.Println("[handle]:", r.URL.EscapedPath())
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodPut:
		{
			b, _ := ioutil.ReadAll(r.Body)
			if len(b) != 0 {
				err := handler.Service.Set(key, b)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
	case http.MethodGet:
		{
			val, err := handler.Service.Get(key)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusNotFound)
			}
			w.Write(val)
		}
	case http.MethodDelete:
		{
			err := handler.Service.Del(key)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusNotFound)
			}
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}
