package ram

type UpdateType byte

const (
	UpdateTypeNone    UpdateType = 0
	UpdateTypeNV      UpdateType = 1
	UpdateTypeOrderly UpdateType = UpdateTypeNV + 1
)
