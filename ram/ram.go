package ram

import "github.com/davidzech/tpm2sim/tpmi"

type RAM struct {
	UpdateNV UpdateType
	// ClearOrderly
	// This flag indicates if the execution of a command should cause the orderly
	// state to be cleared.  This flag is set to FALSE at the beginning of each
	// command in ExecuteCommand() and is checked in ExecuteCommand() after the
	// detailed actions of a command complete but before the check of
	// 'UpdateNV'. If this flag is TRUE, and the orderly state is not
	// SU_NONE_VALUE, then the orderly state in NV memory will be changed to
	// SU_NONE_VALUE or SU_DA_USED_VALUE.
	ClearOrderly bool
	// This location indicates the sequence object handle that holds the DRTM
	// sequence data. When not used, it is set to TPM_RH_UNASSIGNED. A sequence
	// DRTM sequence is started on either _TPM_Init or _TPM_Hash_Start.
	DRTMHandle tpmi.DHObject
}

var UpdateNV UpdateType

// ClearOrderly
// This flag indicates if the execution of a command should cause the orderly
// state to be cleared.  This flag is set to FALSE at the beginning of each
// command in ExecuteCommand() and is checked in ExecuteCommand() after the
// detailed actions of a command complete but before the check of
// 'UpdateNV'. If this flag is TRUE, and the orderly state is not
// SU_NONE_VALUE, then the orderly state in NV memory will be changed to
// SU_NONE_VALUE or SU_DA_USED_VALUE.
//
// TODO(rename this symbol)
var ClearOrderly bool
