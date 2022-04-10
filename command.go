package tpm2sim

import (
	"bytes"
	"fmt"
	"io"

	"github.com/davidzech/tpm2sim/tpm"
	"github.com/davidzech/tpm2sim/tpm/rc"
	"github.com/davidzech/tpm2sim/tpm/st"
	"github.com/davidzech/tpm2sim/tpmi"
)

type Command struct {
	reader     *bytes.Reader
	Size       uint32
	Tag        tpm.ST
	Code       tpm.CC
	Handles    []tpm.Handle // this erases the type, we should probably have this be a slice of interface
	SessionNum int
	AuthSize   int32

	index commandIndex
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
		return nil, rc.ErrCommandSize
	}

	if c.index = getCommandIndex(c.Code); c.index == unimplemented {
		return nil, rc.ErrCommandCode
	}

	// parse Handle Buffer
	err := c.parseHandleBuffer()
	if err != nil {
		return nil, err
	}

	// command.ClearCpRpHashes

	if c.Tag == st.Sessions {
		// calcualte out session buffer size
		if err := tpm.Unmarshal(c.reader, &c.AuthSize); err != nil {
			return nil, err
		}

		if c.AuthSize < 9 || c.AuthSize > int32(c.reader.Len()) {
			return nil, rc.ErrSize
		}

		_, err := c.reader.Seek(int64(c.AuthSize), io.SeekCurrent)
		if err != nil {
			return nil, err
		}

		if err := c.parseSessionBuffer(); err != nil {
			return nil, err
		}
	} else {
		if err := c.checkAuthNoSession(); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Command) Dispatch(response []byte) error {
	panic("not yet")
}

func (c *Command) checkAuthNoSession() error {
	// TODO: #ifdef CC_GetCommandAuditDigest

	for idx := range c.Handles {
		if authRole(c.index, idx) != tpm.AuthNone {
			return rc.ErrAuthMissing
		}
	}

	return nil
}

func (c *Command) parseSessionBuffer() error {
	panic("not yet")
}

func (c *Command) parseHandleBuffer() error {
	// Find the underlying command pointed to by Command.Index
	// This command will have a variadic number of parameters, some of which are specific types oftpm.Handles
	// Find them and unmarshal them in order to validate that they are valid

	descriptor, ok := commandDescriptors[c.index]
	if !ok {
		panic(fmt.Sprintf("bad command index: %v", c.index)) // Index should have been parsed and validated by this point
	}

	handles, err := descriptor.ParameterHandles()
	if err != nil {
		return err
	}

	c.Handles = handles
	return nil
}

var commandDescriptors map[commandIndex]tpmi.CommandDescriptor

// authRole returns the authorization role of a handle
func authRole(cmdIndex commandIndex, handleIndex int) tpm.AuthRole {
	if handleIndex == 0 {
		// TODO: check feature flags to get auth roles for a specific command
	} else if handleIndex == 1 {
		// TODO: check feature flags to get s_commandAttributes[cmdIndex] & HANDLE_2_USER
	}
	return tpm.AuthNone
}

type commandIndex uint16

const (
	unimplemented commandIndex = 0xFFFF
)

func getCommandIndex(cc tpm.CC) commandIndex {
	panic("not yet")
}
