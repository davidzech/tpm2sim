package tpm

const (
	RHFirst       RH = 0x40000000
	RHSRK            = 0x40000000
	RHOwner          = 0x40000001
	RHRevoke         = 0x40000002
	RHTransport      = 0x40000003
	RHOperator       = 0x40000004
	RHAdmin          = 0x40000005
	RHEK             = 0x40000006
	RHNull           = 0x40000007
	RHUnassigned     = 0x40000008
	RSPW             = 0x40000009 // exceptional case of "RS"
	RhLockout        = 0x4000000A
	RHEndorsement    = 0x4000000B
	RHPlatform       = 0x4000000C
	RHPlatformNV     = 0x4000000D
	RHAuth00         = 0x40000010
	RHAuthFF         = 0x4000010F
	RHAct0           = 0x40000110
	RHActF           = 0x4000011F
	RHLast           = 0x4000011F
)
