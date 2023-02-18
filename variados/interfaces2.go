package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	payload := []byte("hello high value softwware engineer")
	hasAndBroadcast(bytes.NewReader(payload))
}

type hashReader struct {
	bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		buf:    bytes.NewBuffer(b),
		Reader: *bytes.NewReader(b),
	}
}

func hasAndBroadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	hash := sha1.Sum(b)
	fmt.Println(hex.EncodeToString(hash[:]))

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of the bytes : ", string(b))
	return nil
}
