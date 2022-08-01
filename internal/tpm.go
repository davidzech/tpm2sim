package internal

type TPM struct {
	nv   *NVRAM
	time *Time

	plat Platform
	g    Globals
}

func (t *TPM) ExecCommand(request []byte) (response []byte) {
	t.g.UpdateNV = UpdateTypeNone
	t.g.ClearOrderly = false

	if t.g.InFailureMode {
		return t.FailureMode(request)
	}

}

func (t *TPM) Started() bool {
	panic("not yet")
}
