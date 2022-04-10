package tpm

import (
	"encoding/binary"
	"io"

	"github.com/davidzech/tpm2sim/tpm/cc"
	"github.com/davidzech/tpm2sim/tpm/hc"
	"github.com/davidzech/tpm2sim/tpm/ht"
	"github.com/davidzech/tpm2sim/tpm/rc"
	"github.com/davidzech/tpm2sim/tpm/rh"
	"github.com/davidzech/tpm2sim/tpm/st"
)

type Handle uint32

func (h Handle) Type() HT {
	return HT((h & hc.HRRangeMask) >> hc.HRShift)
}

type AuthRole uint32

const (
	AuthNone AuthRole = iota
	AuthUser
	AuthAdmin
	AuthDup
)

type CC = cc.CC
type RC = rc.RC
type ST = st.ST
type HC = hc.HC
type HT = ht.HT
type RH = rh.RH

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
