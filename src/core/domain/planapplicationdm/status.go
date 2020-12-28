package planapplicationdm

import "golang.org/x/xerrors"

type Status int

const (
	Applying Status = iota + 1
	Contracted
	Rejected
)

const (
	StatusApplyingStr   string = "applying"
	StatusContractedStr string = "publised"
	StatusRejectedStr   string = "closed"
)

func NewStatus() Status {
	return Status(Applying.Value())
}

func NewStatusForUpdate(status int) (Status, error) {
	if status == int(Applying) {
		return Status(0), xerrors.Errorf("status must not be %d", Applying)
	}

	if status != int(Contracted) && status != int(Rejected) {
		return Status(0), xerrors.Errorf("status must be %d or %d", Contracted, Rejected)
	}

	return Status(status), nil
}

func (s Status) Value() int {
	return int(s)
}

func (s Status) String() string {
	switch s {
	case Applying:
		return StatusApplyingStr
	case Contracted:
		return StatusContractedStr
	case Rejected:
		return StatusRejectedStr
	default:
		return "Undefined"
	}
}

func (s Status) Equals(s2 Status) bool {
	return s == s2
}
