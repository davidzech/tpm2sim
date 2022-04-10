package tpmi

import "github.com/davidzech/tpm2sim/tpm"

type CommandDescriptor interface {
	Call(params ...any) ([]any, error)
	ParameterHandles() ([]tpm.Handle, error)
	InParams() []tpm.Unmarshaler
	OutParams() []tpm.Marshaler
}
