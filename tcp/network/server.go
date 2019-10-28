package network

import (
	"github.com/ShiinaOrez/GoDistributedCache/tcp/cache"
	"github.com/ShiinaOrez/GoDistributedCache/tcp/constvar"
	"net"
	"fmt"
	"log"
	"bufio"
	"io"
)

type Server struct {
	Cache *cache.Service
	Port  string
}

func StartWithService(service *cache.Service) {
    server := &Server {
		Cache: service,
		Port: constvar.Port,
	}
	server.Listen()
}

func (server *Server) Listen() {
	fmt.Println("[start]: listening...")
	listener, err := net.Listen("tcp", server.Port)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go server.Process(conn)
	}
}

func (server *Server) Process(conn net.Conn) {
	defer conn.Close()
	
	reader := bufio.NewReader(conn)
	for {
		op, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("Close connection cause of:", err)
			}
			return
		}
		var res []byte = nil
		switch op {
		case 'G':
			res, err = server.GetFromReader(reader)
		case 'S':
			err = server.SetFromReader(reader)
		case 'D':
			err = server.DelFromReader(reader)
		default: {
			log.Println("Close connection cause of invalid operation!")
			return
		}
		}
		var sErr error = nil
		if err != nil {
			sErr = sendError(conn, err)
		} else {
			sErr = sendResponse(conn, res)
		}
		if sErr != nil {
			log.Println("Close connection cause of:", sErr)
			return
		}
	}
	return
}