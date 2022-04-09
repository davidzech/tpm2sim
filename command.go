package tpm2sim

import (
	"bytes"
	"fmt"
	"io"

	"github.com/davidzech/tpm2sim/tpm"
	"github.com/davidzech/tpm2sim/tpmi"
)

type Command struct {
	reader  io.Reader
	Size    uint32
	Tag     tpm.ST
	Code    tpm.CC
	Index   tpm.CommandIndex
	Handles []tpm.Handle // this erases the type, we should probably have this be a slice of interface
}

func NewCommand(buffer []byte) (*Command, error) {
	c := &Command{
		reader:  bytes.NewReader(buffer),
		Handles: make([]tpm.Handle, 0, 3),
	}

	if err := tpm.Unmarshal(c.reader, (*tpmi.STCommandTag)(&c.Tag)); err != nil {
		return nil, err
	}

	if err := tpm.Unmarshal(c.reader, &c.Size); err != nil {
		return nil, err
	}

	if c.Size != uint32(len(buffer)) || c.Size > tpm.MaxCommandSize {
		return nil, tpm.ErrCommandSize
	}

	c.Index = c.Code.CommandIndex()

	if c.Index == tpm.Unimplemented {
		return nil, tpm.ErrCommandCode
	}

	// parse Handle Buffer
	if err := c.parseHandleBuffer(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Command) parseHandleBuffer() error {
	// Find the underlying command pointed to by Command.Index
	// This command will have a variadic number of parameters, some of which are tpm.Handles
	// Find them and unmarshal them in order to validate that they are valid

	descriptor, ok := commandDescriptors[c.Index]
	if !ok {
		panic(fmt.Sprintf("bad command index: %v", c.Index)) // Index should have been parsed and validated by this point
	}

	for _, param := range descriptor.InParams() {
		handleParam, ok := param.(tpm.HandleUmarshaler)
		if ok {

		}
	}

}

var commandDescriptors map[tpm.CommandIndex]tpmi.CommandDescriptor
