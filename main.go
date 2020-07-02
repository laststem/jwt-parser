package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	
	"github.com/laststem/jwt-parser/internal"
)

const (
	printUsage    = 1
	printAll = 2
)

const usage = `Usage:
  jwt-parser <token>
    Ex)
	    jwt-parser eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
      {
          "sub": "",
          "name": "",
          "iat": 0
      }
	jwt-parser <token> <key> <key> <key> ...
    Ex)
      jwt-parser eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c sub
      1234567890"`

func main() {
	switch len(os.Args) {
	case printUsage:
		fmt.Println(usage)
	case printAll:
		token := parse(os.Args[1])
		data := marshal(token.Payload)
		print(data)
	default:
		token := parse(os.Args[1])
		payload := goDeep(token.Payload, os.Args[2:])
		data := marshal(payload)
		print(data)
	}
}

func parse(raw string) *internal.Token {
	token, err := internal.Parse(raw)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func goDeep(payload map[string]interface{}, key []string) interface{} {
	if len(key) == 1 {
		return payload[key[0]]
	}
	
	remain, ok := payload[key[0]]
	if !ok {
		log.Fatalf("not found key: %v", key)
	}
	
	if p, ok := remain.(map[string]interface{}); ok{
		return goDeep(p, key[1:])
	}
	
	log.Fatalf("not found key: %v", key)
	return nil
}

func marshal(payload interface{}) []byte {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func print(data []byte) {
	data, err := makePretty(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func makePretty(data []byte) ([]byte, error) {
	var out bytes.Buffer
	if err := json.Indent(&out, data, "", "    "); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
