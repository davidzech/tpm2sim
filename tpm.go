package tpm2sim

import (
	"github.com/davidzech/tpm2sim/internal"
	"github.com/davidzech/tpm2sim/internal/nvram"
)

type Simulator struct {
	nvram *nvram.NVRAM

	updateNV      internal.UpdateType
	clearOrderly  bool
	inFailureMode bool
}

func (s *Simulator) ExecCommand(request []byte) (response []byte) {
	s.updateNV = internal.UpdateTypeNone
	s.clearOrderly = false
	if s.inFailureMode {
		return s.failureMode(request)
	}

	s.nvram.CheckState()

}

func (s *Simulator) failureMode(request []byte) (response []byte) {
	panic("not yet")
}
