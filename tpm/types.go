package tpm

import (
	"encoding/binary"
	"io"
)

type RC uint32

type rcErr RC

func (rcErr) Error() string {
	return ""
}

func (rc RC) ToError() error {
	if rc != Success {
		return nil
	}
	return (rcErr(rc))
}

type Handle uint32

type RH Handle

type ST uint16

type Object struct {
	Foo any
}

func (*Object) isObject() {}

type SequenceObject struct {
	Bar any
}

func (*SequenceObject) isObject() {}

type HashObject = SequenceObject

type AnyObject interface {
	isObject()
}

func HandleToObject(handle Handle) AnyObject {
	panic("not yet")
}

type Unmarshaler interface {
	Unmarshal(r io.Reader) error
}

func Unmarshal(r io.Reader, v any) error {
	if unmarshaler, ok := v.(Unmarshaler); ok {
		return unmarshaler.Unmarshal(r)
	}

	return binary.Read(r, binary.BigEndian, v)
}
