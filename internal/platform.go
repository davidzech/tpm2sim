package internal

type Platform interface {
	NVAvailable() TPM_RC
	TimerRead() uint64
	TimerWasStopped() bool
}
