package main

import (
	"bytes"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ShiinaOrez/GoDistributedCache/http-rest/cache"
	"github.com/ShiinaOrez/GoDistributedCache/http-rest/constvar"
	"github.com/ShiinaOrez/GoDistributedCache/http-rest/rest/handler"
)

type pair struct {
	key string
	val string
}

var testSlice []pair

func generateRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func generateTestData() {
	for i := 0; i < 10000; i++ {
		testSlice = append(testSlice, pair{
			key: generateRandomString(6),
			val: generateRandomString(6),
		})
	}
}

func TestCache(t *testing.T) {
	cacheService := cache.NewCacheService()
	handler.SetCacheService(cacheService)
	baseURL := "http://localhost" + constvar.Port + "/cache/"
	cacheServer := httptest.NewUnstartedServer(&handler.CacheHandler)
	cacheServer.URL = baseURL
	defer cacheServer.Close()

	generateTestData()
	for _, pair := range testSlice {
		request := httptest.NewRequest("PUT", baseURL+pair.key, bytes.NewReader([]byte(pair.val)))
		responseRecorder := httptest.NewRecorder()
		handler.CacheHandler.ServeHTTP(responseRecorder, request)
	}
	for _, pair := range testSlice {
		request := httptest.NewRequest("GET", baseURL+pair.key, nil)
		responseRecorder := httptest.NewRecorder()
		handler.CacheHandler.ServeHTTP(responseRecorder, request)
	}
	for _, pair := range testSlice {
		request := httptest.NewRequest("DELETE", baseURL+pair.key, nil)
		responseRecorder := httptest.NewRecorder()
		handler.CacheHandler.ServeHTTP(responseRecorder, request)
	}
}
