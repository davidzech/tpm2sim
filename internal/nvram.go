package internal

type NVRAM struct {
	status TPM_RC
	plat   Platform
}

func (n *NVRAM) CheckState() {
	rc := n.plat.NVAvailable()
	switch rc {
	case 0:
		n.status = TPM_RC_SUCCESS

	case 1:
		n.status = TPM_RC_NV_UNAVAILABLE

	default:
		n.status = TPM_RC_NV_RATE
	}
}

func (n *NVRAM) Available() bool {
	return n.status == TPM_RC_SUCCESS
}
