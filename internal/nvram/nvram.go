package nvram

import "github.com/davidzech/tpm2sim/internal/tpm"

type NVRAM struct {
	evictNVEnd Ref
	status     tpm.RC
}

type Ref uint32

func (n *NVRAM) CheckState() {

	avail := 0 // avail = plat_NVAvailable
	switch avail {
	case 0:
		n.status = tpm.RCSuccess

	case 1:
		n.status = tpm.RCNVUnavailable

	default:
		n.status = tpm.RCNVUnavailable
	}

}

func (n *NVRAM) Status() int
