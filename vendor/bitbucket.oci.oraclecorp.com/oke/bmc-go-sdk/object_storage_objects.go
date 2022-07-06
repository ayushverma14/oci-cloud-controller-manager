// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"time"
)

type ObjectSummary struct {
	Name        string    `json:"name"`
	Size        uint64    `json:"size"`
	MD5         string    `json:"md5"`
	TimeCreated time.Time `json:"timeCreated"`
}

type ListObjects struct {
	OPCClientRequestIDUnmarshaller
	OPCRequestIDUnmarshaller
	Objects       []ObjectSummary `json:"objects"`
	Prefixes      []string        `json:"prefixes"`
	NextStartWith string          `json:"nextStartWith"`
}

type HeadObject struct {
	ContentUnmarshaller
	ETagUnmarshaller
	LastModifiedUnmarshaller
	MetadataUnmarshaller
	OPCClientRequestIDUnmarshaller
	OPCRequestIDUnmarshaller
	ID        string
	Bucket    string
	Namespace Namespace
}

// Object is an item stored in ObjectStorage
type Object struct {
	HeadObject
	Size         uint64
	TraceID      string
	Body         []byte
	BodyAsStream io.ReadCloser
}

func (g *Object) SetBody(b []byte, toBeFilled interface{}) error {
	rv := reflect.ValueOf(toBeFilled)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("Value passed to unmarshal is not a pointer")
	}
	if po, ok := toBeFilled.(*Object); ok {
		po.Body = b
	} else {
		return errors.New("Value passed in was not an Object")
	}
	return nil
}

func (g *Object) SetBodyAsStream(b io.ReadCloser, toBeFilled interface{}) error {
	rv := reflect.ValueOf(toBeFilled)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("Value passed to unmarshal is not a pointer")
	}
	if po, ok := toBeFilled.(*Object); ok {
		po.BodyAsStream = b
	} else {
		return errors.New("Value passed in was not an Object")
	}
	return nil
}

type DeleteObject struct {
	OPCRequestIDUnmarshaller
	OPCClientRequestIDUnmarshaller
	LastModifiedUnmarshaller
}

// ListObjects lists objects
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/ListObjects/ListObjects
func (c *Client) ListObjects(namespace Namespace, bucket string, opts *ListObjectsOptions) (objects *ListObjects, e error) {
	details := &requestDetails{
		ids: urlParts{
			namespace,
			resourceBuckets,
			bucket,
			resourceObjects,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	objects = &ListObjects{}
	e = resp.unmarshal(objects)
	return
}

// GetObject fetches an object from object storage
// Returns object of []byte type in Object.Body
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/methods/GetObject
func (c *Client) GetObject(
	namespace Namespace,
	bucketName string,
	objectName string,
	opts *GetObjectOptions,
) (object *Object, e error) {
	details := &requestDetails{
		ids: urlParts{
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	object = &Object{}
	e = resp.unmarshal(object)
	object.Namespace = namespace
	object.Bucket = bucketName
	object.ID = objectName
	return
}

// GetObjectAsStream fetches an object from object storage
// Returns object of io.ReadCloser type in Object.BodyAsStream with the use of returnRespBodyAsStream set to true in requestDetails
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/methods/GetObject
func (c *Client) GetObjectAsStream(
	namespace Namespace,
	bucketName string,
	objectName string,
	opts *GetObjectOptions,
) (object *Object, e error) {
	details := &requestDetails{
		ids: urlParts{
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional:               opts,
		returnRespBodyAsStream: true,
	}

	var resp *response
	if resp, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	object = &Object{}
	e = resp.unmarshal(object)
	object.Namespace = namespace
	object.Bucket = bucketName
	object.ID = objectName
	return
}

// DeleteObject deletes an object from object storage
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/methods/DeleteObject
func (c *Client) DeleteObject(
	namespace Namespace,
	bucketName string,
	objectName string,
	opts *DeleteObjectOptions,
) (object *DeleteObject, e error) {

	details := &requestDetails{
		ids: urlParts{
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.objectStorageApi.request(http.MethodDelete, details); e != nil {
		return
	}

	object = &DeleteObject{}
	e = resp.unmarshal(object)
	return
}

// HeadObject fetches the user defined metadata for an object
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/methods/HeadObject
func (c *Client) HeadObject(
	namespace Namespace,
	bucketName string,
	objectName string,
	opts *HeadObjectOptions,
) (headObject *HeadObject, e error) {

	details := &requestDetails{
		ids: urlParts{
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.objectStorageApi.request(http.MethodHead, details); e != nil {
		return
	}

	headObject = &HeadObject{}
	e = resp.unmarshal(headObject)
	headObject.Namespace = namespace
	headObject.Bucket = bucketName
	headObject.ID = objectName
	return
}

// PutObject updates an object in object storage
// Uses object of []byte type
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/methods/PutObject
func (c *Client) PutObject(
	namespace Namespace,
	bucketName string,
	objectName string,
	content []byte,
	opts *PutObjectOptions,
) (object *Object, e error) {

	required := struct {
		bodyRequirement
		ContentLength uint64 `header:"Content-Length" json:"-" url:"-"`
	}{
		ContentLength: uint64(len(content)),
	}
	required.Body = content

	details := &requestDetails{
		ids: urlParts{
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.objectStorageApi.request(http.MethodPut, details); e != nil {
		return
	}

	object = &Object{}
	e = resp.unmarshal(object)
	object.Namespace = namespace
	object.Bucket = bucketName
	object.ID = objectName
	object.Body = content
	return
}

// PutObjectAsStream updates an object in object storage
// Uses object of io.ReadSeeker type
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/methods/PutObject
func (c *Client) PutObjectAsStream(
	namespace Namespace,
	bucketName string,
	objectName string,
	content io.ReadSeeker,
	opts *PutObjectOptions,
) (object *Object, e error) {

	required := struct {
		bodyRequirement
	}{}
	required.Body = content

	details := &requestDetails{
		ids: urlParts{
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.objectStorageApi.request(http.MethodPut, details); e != nil {
		return
	}

	object = &Object{}
	e = resp.unmarshal(object)
	object.Namespace = namespace
	object.Bucket = bucketName
	object.ID = objectName
	// The body is a stream and it's already closed after request is completed, so just set this to nil
	object.Body = nil
	return
}
