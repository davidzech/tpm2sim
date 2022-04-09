package tpmi

import (
	"io"

	"github.com/davidzech/tpm2sim/tpm"
)

type DHObject tpm.Handle

type STCommandTag tpm.ST

func (s *STCommandTag) Unmarshal(r io.Reader) error {
	return (*tpm.ST)(s).Unmarshal(r)
}

type Platform tpm.Handle
