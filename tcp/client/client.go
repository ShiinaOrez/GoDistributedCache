package main

import (
	"net"
	"flag"
	"log"
	"os"
	"fmt"
	"strconv"
	"bufio"
	"github.com/ShiinaOrez/GoDistributedCache/tcp/constvar"
)

func generateBytes(str string) []byte {
	length := len(str)
	lenStr := strconv.Itoa(length)
	return append([]byte(lenStr+" "), []byte(str)...)
}

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

	bytes := []byte{}
	switch (*cmd) {
	case "get": {
		if (*key) == "" {
			log.Println("Get must with key.")
			os.Exit(1)
		}
		bytes = append([]byte{'G'}, generateBytes((*key))...)
	}
	case "set": {
		if (*key) == "" || (*val) == "" {
			log.Println("Set must with key and value.")
			os.Exit(1)
		}
		bytes = append([]byte{'S'}, append(generateBytes((*key)), generateBytes(*val)...)...)
	}
	case "del": {
		if (*key) == "" {
			log.Println("Del must with key.")
			os.Exit(1)
		}
		bytes = append([]byte{'G'}, generateBytes((*key))...)
	}
	}
	conn.Write(bytes)
	reader := bufio.NewReader(conn)
	res, err := reader.ReadString(' ')
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