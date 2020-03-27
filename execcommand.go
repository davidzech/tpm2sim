package tpm2sim

import (
	"github.com/davidzech/tpm2sim/fail"
	"github.com/davidzech/tpm2sim/nv"
	"github.com/davidzech/tpm2sim/object"
	"github.com/davidzech/tpm2sim/ram"
	"github.com/davidzech/tpm2sim/time"
	"github.com/davidzech/tpm2sim/tpm"
)

type Command struct {
	ParameterBuffer []byte
}

func ExecuteCommand(request []byte) (response []byte) {

	var commandSize uint32
	_ = commandSize // silence
	var command Command

	var result tpm.RC
	_ = result // silence

	ram.UpdateNV = ram.UpdateTypeNone
	ram.ClearOrderly = false

	if fail.InFailureMode {

		return fail.TpmFailureMode(request)
	}

	// Query platform to get the NV state.  The result state is saved internally
	// and will be reported by NvIsAvailable(). The reference code requires that
	// accessibility of NV does not change during the execution of a command.
	// Specifically, if NV is available when the command execution starts and then
	// is not available later when it is necessary to write to NV, then the TPM
	// will go into failure mode.
	nv.CheckState()

	// Due to the limitations of the simulation, TPM clock must be explicitly
	// synchronized with the system clock whenever a command is received.
	// This function call is not necessary in a hardware TPM. However, taking
	// a snapshot of the hardware timer at the beginning of the command allows
	// the time value to be consistent for the duration of the command execution.
	time.UpdateToCurrent()

	// Any command through this function will unceremoniously end the
	// _TPM_Hash_Data/_TPM_Hash_End sequence.
	// TODO: rename these symbols

	if ram.DRTMHandle != tpm.RHUnassigned {
		object.TerminateEvent()
	}
	// if(g_DRTMHandle != TPM_RH_UNASSIGNED)
	// ObjectTerminateEvent()

	command.ParameterBuffer = request

	// Parse command header: tag, commandSize and command.code.
	// First parse the tag. The unmarshaling routine will validate
	// that it is either TPM_ST_SESSIONS or TPM_ST_NO_SESSIONS.
	return
}
