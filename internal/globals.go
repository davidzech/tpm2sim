package internal

type Globals struct {
	UpdateNV         UpdateType
	PowerWasLost     bool
	ClearOrderly     bool
	PrevOrderlyState bool
	NVOK             bool
	InFailureMode    bool
}

type ORDERLY_DATA struct {
	clock uint64
}
