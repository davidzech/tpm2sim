package tpmi

import (
	"encoding/binary"
	"io"

	"github.com/davidzech/tpm2sim/tpm"
)

type DHObject tpm.Handle

type STCommandTag tpm.ST

func (s *STCommandTag) Unmarshal(r io.Reader) error {
	// read uint16 out of buf
	var u16 *uint16 = (*uint16)(s)
	err := binary.Read(r, binary.BigEndian, u16)
	if err != nil {
		return err
	}

	return nil
}
