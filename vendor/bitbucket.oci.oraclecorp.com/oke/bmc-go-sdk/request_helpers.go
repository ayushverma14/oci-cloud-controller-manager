// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type authenticationInfo struct {
	privateRSAKey  *rsa.PrivateKey
	tenancyOCID    string
	userOCID       string
	keyFingerPrint string
}

func (a *authenticationInfo) getKeyID() string {
	return fmt.Sprintf("%s/%s/%s", a.tenancyOCID, a.userOCID, a.keyFingerPrint)
}

func getErrorFromResponse(body io.Reader, resp *http.Response) (e error) {
	apiError := Error{}

	if opcRequestID := resp.Header.Get(headerOPCRequestID); opcRequestID != "" {
		apiError.OPCRequestID = opcRequestID
	}

	if resp.Header.Get("content-type") != "application/json" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(body)
		apiError.Message = buf.String()
	} else {
		decoder := json.NewDecoder(body)
		if e = decoder.Decode(&apiError); e != nil {
			buf := new(bytes.Buffer)
			buf.ReadFrom(decoder.Buffered())
			apiError.Message = buf.String()
		}
	}
	apiError.Status = strconv.Itoa(resp.StatusCode)

	return &apiError
}

func createAuthorizationHeader(request *http.Request, auth *authenticationInfo, userAgent string, body interface{}) (e error) {
	addRequiredRequestHeaders(request, userAgent, body)
	var sig string

	if sig, e = computeSignature(request, auth.privateRSAKey); e != nil {
		return
	}

	signedHeaders := getSigningHeaders(request.Method)
	headers := concatenateHeaders(signedHeaders)

	authValue := fmt.Sprintf("Signature headers=\"%s\",keyId=\"%s\",algorithm=\"rsa-sha256\",signature=\"%s\"", headers, auth.getKeyID(), sig)

	request.Header.Add("authorization", authValue)

	return
}

func concatenateHeaders(headers []string) (concatenated string) {

	for _, header := range headers {
		if len(concatenated) > 0 {
			concatenated += " "
		}
		concatenated += header
	}

	return
}

func getSigningHeaders(method string) []string {
	result := []string{
		"date",
		"(request-target)",
	}

	if method == http.MethodPost || method == http.MethodPut {
		result = append(result, "content-length", "content-type", "x-content-sha256")
	}

	return result
}

func computeSignature(request *http.Request, privateKey *rsa.PrivateKey) (sig string, e error) {
	signingString := getSigningString(request)
	hasher := sha256.New()
	hasher.Write([]byte(signingString))
	hashed := hasher.Sum(nil)
	var unencodedSig []byte
	unencodedSig, e = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if e != nil {
		return
	}

	sig = base64.StdEncoding.EncodeToString(unencodedSig)

	return

}

func getSigningString(request *http.Request) string {
	signingHeaders := getSigningHeaders(request.Method)
	signingString := ""
	for _, header := range signingHeaders {
		if signingString != "" {
			signingString += "\n"
		}

		if header == "(request-target)" {
			signingString += fmt.Sprintf("%s: %s", header, getRequestTarget(request))
		} else {
			signingString += fmt.Sprintf("%s: %s", header, request.Header.Get(header))
		}
	}

	return signingString

}

func getRequestTarget(request *http.Request) string {
	lowercaseMethod := strings.ToLower(request.Method)
	return fmt.Sprintf("%s %s", lowercaseMethod, request.URL.RequestURI())
}

func addIfNotPresent(dest *http.Header, key, value string) {
	if dest.Get(key) == "" {
		dest.Set(key, value)
	}
}

func getBodyHash(body []byte) string {
	hash := sha256.Sum256(body)
	return base64.StdEncoding.EncodeToString(hash[:])
}

// This will return the length of the io.ReadSeeker content and its SHA-256 hash computation.
// Once the hash computation is completed, the io.ReadSeeker's offset is reset back to beginning
// of the stream so the next operation on it will start from there.
func getBodyStreamLengthAndHash(body io.ReadSeeker) (int64, string) {
	var e error
	var length int64

	hash := sha256.New()
	length, e = io.Copy(hash, body)
	if e != nil {
		panic(e.Error())
	}
	_, _ = body.Seek(0, io.SeekStart)
	if e != nil {
		panic(e.Error())
	}
	return length, hex.EncodeToString(hash.Sum(nil))
}

func addRequiredRequestHeaders(request *http.Request, userAgent string, body interface{}) {
	addIfNotPresent(&request.Header, "content-type", "application/json")
	addIfNotPresent(&request.Header, "date", time.Now().UTC().Format(http.TimeFormat))
	if userAgent == "" {
		addIfNotPresent(&request.Header, "User-Agent", fmt.Sprintf("baremetal-sdk-go-v%s", SDKVersion))
	} else {
		addIfNotPresent(&request.Header, "User-Agent", userAgent)
	}
	addIfNotPresent(&request.Header, "accept", "*/*")

	if request.Method == http.MethodPost || request.Method == http.MethodPut {
		var bodyHash string
		switch bodyValue := body.(type) {
		case []byte:
			bodyHash = getBodyHash(bodyValue)
		case io.ReadSeeker:
			request.ContentLength, bodyHash = getBodyStreamLengthAndHash(bodyValue)
		}
		addIfNotPresent(&request.Header, "content-length", strconv.FormatInt(request.ContentLength, 10))

		if request.ContentLength > 0 {
			addIfNotPresent(&request.Header, "x-content-sha256", bodyHash)
		}

	}
}