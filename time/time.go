package time

import (
	"time"

	"github.com/davidzech/tpm2sim/nv"
)

type Clock struct {
	nv       *nv.NV
	power    interface{ Started() bool }
	platform interface {
		TimerWasStopped() bool
		TimerRead() time.Time
	}
	time time.Time
}

func (c *Clock) UpdateToCurrent() {
	if !c.nv.Available() || !c.power.Started() {
		return
	}

	c.TimeUpdate()
}

func (c *Clock) TimeUpdate() {
	var elapsed time.Duration

	if c.platform.TimerWasStopped() {
		panic("TimeNewEpoch()")
		// TimeNewEpoch()
	}

	elapsed = c.platform.TimerRead().Sub(c.time)

	c.time.Add(time.Duration(elapsed))

	panic("TimeClockUpdate(go.clock + elapsed)")

	panic("DASelfHeal()")
}
