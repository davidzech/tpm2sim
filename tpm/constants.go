package tpm

import "encoding/binary"

const (
	MaxCommandSize = 4096
)

var (
	STDResponseHeader = binary.Size(ST(0)) + binary.Size(uint32(0)) + binary.Size(RC(0))
)
