package rc

import "fmt"

type Error RC

func (r Error) Error() string {
	return fmt.Sprintf("tpm rc: %d", int(r))
}

func (r Error) RC() RC {
	return RC(r)
}

func (rc RC) AsErr() error {
	if rc != Success {
		return nil
	}
	return (Error(rc))
}
