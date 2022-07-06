// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

package resourceprincipals

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/oci-go-sdk/v49/common"
)

type securityToken interface {
	fmt.Stringer
	Valid() bool
	ExpiresAt() time.Time
}

type token struct {
	tokenString string
	jwtToken    *JwtToken
}

func newToken(tokenString string) (newToken securityToken, err error) {
	var jwtToken *JwtToken
	if jwtToken, err = ParseJwt(tokenString); err != nil {
		return nil, fmt.Errorf("failed to parse the token string \"%s\": %s", tokenString, err.Error())
	}
	return &token{tokenString, jwtToken}, nil
}

func (t *token) String() string {
	return t.tokenString
}

func (t *token) Valid() bool {
	return !t.jwtToken.Expired()
}

func (t *token) ExpiresAt() time.Time {
	return t.jwtToken.Expiry()
}

// JwtToken Representation of a JSON web token
type JwtToken struct {
	raw     string
	header  map[string]interface{}
	payload map[string]interface{}
}

const bufferTimeBeforeTokenExpiration = 5 * time.Minute

// Expiry Gets the token xpiry time
func (t *JwtToken) Expiry() time.Time {
	exp := int64(t.payload["exp"].(float64))
	return time.Unix(exp, 0)
}

// Expired Determines if the token is expired based on the current time
func (t *JwtToken) Expired() bool {
	exp := int64(t.payload["exp"].(float64))
	expTime := time.Unix(exp, 0)
	expired := exp <= time.Now().Unix()+int64(bufferTimeBeforeTokenExpiration.Seconds())
	if expired {
		common.Debugf("Token expires at:  %v, currently expired due to bufferTime: %v", expTime.Format("15:04:05.000"), expired)
	}
	return expired
}

// Raw Returns the raw JWT string
func (t *JwtToken) Raw() string {
	return t.raw
}

// Header returns the jwt header
func (t *JwtToken) Header() map[string]interface{} {
	return t.header
}

// Payload returns the jwt payload
func (t *JwtToken) Payload() map[string]interface{} {
	return t.payload
}

// ParseJwt Converts a string representation of a JWT to a JWT struct
func ParseJwt(tokenString string) (*JwtToken, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("the given token string contains an invalid number of parts")
	}

	token := &JwtToken{raw: tokenString}
	var err error

	// Parse Header part
	var headerBytes []byte
	if headerBytes, err = decodePart(parts[0]); err != nil {
		return nil, fmt.Errorf("failed to decode the header bytes: %s", err.Error())
	}
	if err = json.Unmarshal(headerBytes, &token.header); err != nil {
		return nil, err
	}

	// Parse Payload part
	var payloadBytes []byte
	if payloadBytes, err = decodePart(parts[1]); err != nil {
		return nil, fmt.Errorf("failed to decode the payload bytes: %s", err.Error())
	}
	decoder := json.NewDecoder(bytes.NewBuffer(payloadBytes))
	if err = decoder.Decode(&token.payload); err != nil {
		return nil, fmt.Errorf("failed to decode the payload json: %s", err.Error())
	}

	return token, nil
}

func decodePart(partString string) ([]byte, error) {
	if l := len(partString) % 4; 0 < l {
		partString += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(partString)
}
