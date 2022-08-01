package tpm2sim

import (
	"github.com/davidzech/tpm2sim/internal"
	"github.com/davidzech/tpm2sim/internal/globals"
	"github.com/davidzech/tpm2sim/internal/nvram"
	"github.com/davidzech/tpm2sim/internal/time"
	"github.com/davidzech/tpm2sim/internal/tpm"
)

type Simulator struct {
	nvram *nvram.NVRAM
	time  *time.Clock

	g globals.Globals
}

func (s *Simulator) ExecCommand(request []byte) (response []byte) {
	s.g.UpdateNV = internal.UpdateTypeNone
	s.g.ClearOrderly = false
	if s.g.InFailureMode {
		return s.failureMode(request)
	}

	s.nvram.CheckState()

	s.time.UpdateToCurrent()

	if s.g.DRTMHandle != tpm.RHUnassigned {

	}
}

func (s *Simulator) failureMode(request []byte) (response []byte) {
	panic("not yet")
}
