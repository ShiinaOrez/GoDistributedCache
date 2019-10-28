package network

import (
	"bufio"
)

func (server *Server) GetFromReader(reader *bufio.Reader) ([]byte, error) {
	key, err := getKey(reader)
	if err != nil {
		return nil, err
	}
	val, err := server.get(key)
	if err != nil {
		return nil, err
	}
    return val, nil
}

func (server *Server) SetFromReader(reader *bufio.Reader) error {
	key, err := getKey(reader)
	if err != nil {
		return err
	}
	val, err := getVal(reader)
	if err != nil {
		return err
	}
	err = server.set(key, val)
	if err != nil {
		return err
	}
    return nil
}

func (server *Server) DelFromReader(reader *bufio.Reader) error {
	key, err := getKey(reader)
	if err != nil {
		return err
	}
	err = server.del(key)
	if err != nil {
		return err
	}
    return nil
}

func (server *Server) get(key string) ([]byte, error) {
	return server.Cache.Get(key)
}

func (server *Server) set(key string, val []byte) error {
	return server.Cache.Set(key, val)
}

func (server *Server) del(key string) error {
	return server.Cache.Del(key)
}