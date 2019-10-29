package main

import (
	"math/rand"
	"testing"
	"time"
	"net"
	"log"
	"os"
	"github.com/ShiinaOrez/GoDistributedCache/tcp/cache"
	"github.com/ShiinaOrez/GoDistributedCache/tcp/constvar"
	"github.com/ShiinaOrez/GoDistributedCache/tcp/network"
	"github.com/ShiinaOrez/GoDistributedCache/tcp/client/tools"
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

func TestTCP(t *testing.T) {
	cacheService := cache.NewService()
	network.StartWithService(cacheService)

	server, err := net.ResolveTCPAddr("tcp", "localhost"+constvar.Port)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	generateTestData()
	for _, pair := range testSlice {
		tools.Request(conn, "localhost", "set", pair.key, pair.val)
		tools.Request(conn, "localhost", "get", pair.key, "")
		tools.Request(conn, "localhost", "del", pair.key, "")
	}
}