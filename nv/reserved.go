package nv

import "github.com/davidzech/tpm2sim/tpm"

type NV struct {
	platform any
	status   tpm.RC
}

func New(platform any) *NV {
	nv := &NV{
		platform: platform,
	}

	nv.CheckState()

	return nv
}

func (nv *NV) CheckState() {
	nvAvailable := tpm.Warn

	nv.status = nvAvailable
}

func (nv *NV) Available() bool {
	return nv.status == tpm.Success
}
