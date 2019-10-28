package network

import (
	"net"
	"fmt"
)

func sendResponse(conn net.Conn, res []byte) error {
	length := fmt.Sprintf("+%d ", len(res))
	_, err := conn.Write(append([]byte(length), res...))
    return err
}

func sendError(conn net.Conn, err error) error {
	length := fmt.Sprintf("-%d ", len(err.Error()))
	_, e := conn.Write(append([]byte(length), []byte(err.Error())...))
	return e
}