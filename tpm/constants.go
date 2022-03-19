package tpm

const (
	First       RH = 0x40000000
	SRK         RH = 0x40000000
	Owner       RH = 0x40000001
	Revoke      RH = 0x40000002
	Transport   RH = 0x40000003
	Operator    RH = 0x40000004
	Admin       RH = 0x40000005
	EK          RH = 0x40000006
	Null        RH = 0x40000007
	Unassigned  RH = 0x40000008
	PW          RH = 0x40000009 // exceptional case of "RS"
	Lockout     RH = 0x4000000A
	Endorsement RH = 0x4000000B
	Platform    RH = 0x4000000C
	PlatformNV  RH = 0x4000000D
	Auth00      RH = 0x40000010
	AuthFF      RH = 0x4000010F
	Act0        RH = 0x40000110
	ActF        RH = 0x4000011F
	Last        RH = 0x4000011F
)

const (
	Success       RC = 0
	Warn          RC = 0x900
	NVRate        RC = Warn + 0x20
	NVUnavailable RC = Warn + 0x23
)
