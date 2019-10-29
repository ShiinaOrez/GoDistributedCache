package network

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func getKey(reader *bufio.Reader) (string, error) {
	length, err := getLen(reader)
	if err != nil {
		return "", err
	}
	key := make([]byte, length)
	_, err = io.ReadFull(reader, key)
	if err != nil {
		return "", err
	}
	return string(key), nil
}

func getVal(reader *bufio.Reader) ([]byte, error) {
	length, err := getLen(reader)
	if err != nil {
		return nil, err
	}
	val := make([]byte, length)
	_, err = io.ReadFull(reader, val)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func getLen(reader *bufio.Reader) (int, error) {
	tmp, err := reader.ReadString(' ')
	if err != nil {
		return 0, err
	}
	length, err := strconv.Atoi(strings.TrimSpace(tmp))
	return length, nil
}
