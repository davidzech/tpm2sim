package internal

type Time struct {
	tpm *TPM

	time uint64
}

func (t *Time) UpdateToCurrent() {
	if !t.tpm.nv.Available() || !t.tpm.Started() {
		return
	}

	t.Update()
}

func (t *Time) Update() {

	if t.tpm.plat.TimerWasStopped() {
		t.NewEpoch()
	}

	elapsed := t.tpm.plat.TimerRead() - t.time

	t.time += elapsed

	panic("t.ClockUpdate(go.clock + elapsed)")

	panic("DASelfHeal()")
}

func (t *Time) NewEpoch() {
	panic("not yet")
}

func (t *Time) ClockUpdate(newTime uint64)
