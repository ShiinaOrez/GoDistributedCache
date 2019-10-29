package tools

import (
	_ "log"
	"net"
	"os"
	"strconv"
	"bufio"
)

func generateBytes(str string) []byte {
	length := len(str)
	lenStr := strconv.Itoa(length)
	return append([]byte(lenStr+" "), []byte(str)...)
}

func Request(conn net.Conn, host, cmd, key, val string) (string, error) {
	bytes := []byte{}
	switch cmd {
	case "get": {
		if key == "" {
			// log.Println("Get must with key.")
			os.Exit(1)
		}
		bytes = append([]byte{'G'}, generateBytes(key)...)
	}
	case "set": {
		if key == "" || val == "" {
			// log.Println("Set must with key and value.")
			os.Exit(1)
		}
		bytes = append([]byte{'S'}, append(generateBytes(key), generateBytes(val)...)...)
	}
	case "del": {
		if key == "" {
			// log.Println("Del must with key.")
			os.Exit(1)
		}
		bytes = append([]byte{'D'}, generateBytes(key)...)
	}
	}
	conn.Write(bytes)
	reader := bufio.NewReader(conn)
	return reader.ReadString(' ')
}