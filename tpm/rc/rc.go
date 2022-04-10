package rc

type RC uint32

const (
	Success       RC = 0
	Ver1          RC = 0x00000100
	CommandSize   RC = Ver1 + 0x42
	Warn          RC = 0x00000900
	NVRate        RC = Warn + 0x20
	NVUnavailable RC = Warn + 0x23
	CommandCode   RC = Ver1 + 0x43
	Initialize    RC = Ver1
	Size          RC = 0x80 + 0x15
	AuthMissing   RC = Ver1 + 0x25
)

var (
	ErrSize        = Size.AsErr()
	ErrCommandSize = CommandSize.AsErr()
	ErrCommandCode = CommandCode.AsErr()
	ErrInitialize  = Initialize.AsErr()
	ErrAuthMissing = AuthMissing.AsErr()
)
