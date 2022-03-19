package tpm2sim

import (
	"bytes"
	"io"

	"github.com/davidzech/tpm2sim/fail"
	"github.com/davidzech/tpm2sim/nv"
	"github.com/davidzech/tpm2sim/ram"
	"github.com/davidzech/tpm2sim/time"
	"github.com/davidzech/tpm2sim/tpm"
	"github.com/davidzech/tpm2sim/tpmi"
)

type Simulator struct {
	nv    *nv.NV
	ram   *ram.RAM
	clock *time.Clock
}

type Command struct {
	io.Reader
	tag  tpmi.STCommandTag
	size uint32
}

func NewCommand(buffer []byte) (*Command, error) {
	c := &Command{
		Reader: bytes.NewReader(buffer),
	}

	if err := tpm.Unmarshal(c, &c.tag); err != nil {
		return nil, err
	}

	if err := tpm.Unmarshal(c, &c.size); err != nil {
		return nil, err
	}

}

// ExecuteCommand
//
// The function performs the following steps.
//
//  a)  Parses the command header from input buffer.
//  b)  Calls ParseHandleBuffer() to parse the handle area of the command.
//  c)  Validates that each of the handles references a loaded entity.
//  d)  Calls ParseSessionBuffer () to:
//      1)  unmarshal and parse the session area;
//      2)  check the authorizations; and
//      3)  when necessary, decrypt a parameter.
//  e)  Calls CommandDispatcher() to:
//      1)  unmarshal the command parameters from the command buffer;
//      2)  call the routine that performs the command actions; and
//      3)  marshal the responses into the response buffer.
//  f)  If any error occurs in any of the steps above create the error response
//      and return.
//  g)  Calls BuildResponseSessions() to:
//      1)  when necessary, encrypt a parameter
//      2)  build the response authorization sessions
//      3)  update the audit sessions and nonces
//  h)  Calls BuildResponseHeader() to complete the construction of the response.
//
// 'responseSize' is set by the caller to the maximum number of bytes available in
// the output buffer. ExecuteCommand will adjust the value and return the number
// of bytes placed in the buffer.
//
// 'response' is also set by the caller to indicate the buffer into which
//  ExecuteCommand is to place the response.
//
//  'request' and 'response' may point to the same buffer
//
// Note: As of February, 2016, the failure processing has been moved to the
// platform-specific code. When the TPM code encounters an unrecoverable failure, it
// will SET g_inFailureMode and call _plat__Fail(). That function should not return
// but may call ExecuteCommand().
//
func (s *Simulator) ExecuteCommand(request []byte) (response []byte) {

	s.ram.UpdateNV = ram.UpdateTypeNone
	s.ram.ClearOrderly = false

	if fail.InFailureMode {

		return fail.TpmFailureMode(request)
	}

	// Query platform to get the NV state.  The result state is saved internally
	// and will be reported by NvIsAvailable(). The reference code requires that
	// accessibility of NV does not change during the execution of a command.
	// Specifically, if NV is available when the command execution starts and then
	// is not available later when it is necessary to write to NV, then the TPM
	// will go into failure mode.
	s.nv.CheckState()

	// Due to the limitations of the simulation, TPM clock must be explicitly
	// synchronized with the system clock whenever a command is received.
	// This function call is not necessary in a hardware TPM. However, taking
	// a snapshot of the hardware timer at the beginning of the command allows
	// the time value to be consistent for the duration of the command execution.
	s.clock.UpdateToCurrent()

	// Any command through this function will unceremoniously end the
	// _TPM_Hash_Data/_TPM_Hash_End sequence.
	// TODO: rename these symbols

	if tpm.RH(s.ram.DRTMHandle) != tpm.Unassigned {
		s.objectTerminateEvent()
	}

	// Parse command header: tag, commandSize and command.code.
	// First parse the tag. The unmarshaling routine will validate
	// that it is either TPM_ST_SESSIONS or TPM_ST_NO_SESSIONS.
	// result = TPMI_ST_COMMAND_TAG_Unmarshal(&command.tag,
	// 	&command.parameterBuffer,
	// 	&command.parameterSize)

	command := NewCommand(request)

	// Parse command header: tag, commandSize and command.code.
	// First parse the tag. The unmarshaling routine will validate
	// that it is either TPM_ST_SESSIONS or TPM_ST_NO_SESSIONS.
	return
}

// ObjectTerminateEvent
// This function is called to close out the event sequence and clean up the hash
// context states.
func (s *Simulator) objectTerminateEvent() {

}
