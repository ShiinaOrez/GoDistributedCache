package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"net"
	"github.com/ShiinaOrez/GoDistributedCache/tcp/client/tools"
	"github.com/ShiinaOrez/GoDistributedCache/tcp/constvar"
)

func main() {
	host := flag.String("h", "localhost", "cache server address.")
	cmd := flag.String("c", "get", "command, must be get/set/del")
	key := flag.String("key", "", "key")
	val := flag.String("value", "", "value")
	flag.Parse()

	server, err := net.ResolveTCPAddr("tcp", (*host)+constvar.Port)
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

	res, err := tools.Request(conn, (*host), (*cmd), (*key), (*val))
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	if res[0] == '-' {
		// error
		fmt.Println(res)
	}
	if res[0] == '+' {
		fmt.Println(res)
	}
	return
}