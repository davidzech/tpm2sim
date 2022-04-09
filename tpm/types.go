package tpm

import (
	"encoding/binary"
	"fmt"
	"io"
)

type CC uint32

func (c CC) CommandIndex() CommandIndex {
	panic("not yet")
}

type CommandIndex uint16

type RC uint32

type RCError RC

func (r RCError) Error() string {
	return fmt.Sprintf("tpm rc: %d", int(r))
}

func (r RCError) RC() RC {
	return RC(r)
}

func (rc RC) AsErr() error {
	if rc != Success {
		return nil
	}
	return (RCError(rc))
}

type Handle uint32

type RH Handle

type ST uint16

func (s *ST) Unmarshal(r io.Reader) error {
	// read uint16 out of buf
	var u16 *uint16 = (*uint16)(s)
	err := binary.Read(r, binary.BigEndian, u16)
	if err != nil {
		return err
	}

	return nil
}

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
	UnmarshalTPM(r io.Reader) error
}

func Unmarshal(r io.Reader, v any) error {
	if unmarshaler, ok := v.(Unmarshaler); ok {
		return unmarshaler.UnmarshalTPM(r)
	}

	return binary.Read(r, binary.BigEndian, v)
}

type Marshaler interface {
	MarshalTPM(io.Writer) error
}
