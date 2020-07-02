package jwt_parser

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type Token struct {
	Payload map[string]interface{}
}

func Parse(raw string) (*Token, error) {
	parts := strings.Split(raw, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token")
	}
	
	payload := parts[1]
	decoded, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}
	
	var m map[string]interface{}
	if err := json.Unmarshal(decoded, &m); err != nil {
		return nil, err
	}
	return &Token{Payload: m}, nil
}