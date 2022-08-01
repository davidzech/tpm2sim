package tpm

type RC uint32

const (
	RCSuccess       RC = 0
	RCWarn          RC = 0x900
	RCNVUnavailable RC = RCWarn + 0x023
	RCNVRate        RC = RCWarn + 0x020
)
