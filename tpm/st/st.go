package st

import (
	"encoding/binary"
	"io"
)

type ST uint16

func (s *ST) Unmarshal(r io.Reader) error {
	// read uint16 out of buf
	var u16 *uint16 = (*uint16)(s)
	err := binary.Read(r, binary.BigEndian, u16)
	if err != nil {
		return err
	}

	return nil
}

const (
	RSPCommand          ST = 0x00c4
	Null                ST = 0x8000
	NoSessions          ST = 0x8001
	Sessions            ST = 0x8002
	AttestNV            ST = 0x8104
	AttestCommandAudit  ST = 0x8015
	AttestSesssionAudit ST = 0x8016
	AttestCertify       ST = 0x8017
	AttestTime          ST = 0x8019
	AttestCreation      ST = 0x801a
	AttestNVDigest      ST = 0x801c
	Creation            ST = 0x8021
	Verified            ST = 0x8022
	AuthSecret          ST = 0x8023
	HashCheck           ST = 0x8024
	AuthSigned          ST = 0x8025
	FUManifest          ST = 0x8029
)
