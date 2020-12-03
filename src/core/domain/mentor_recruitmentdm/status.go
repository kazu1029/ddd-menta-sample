package mentor_recruitmentdm

import "golang.org/x/xerrors"

type Status int

const (
	Draft Status = iota + 1
	Published
	Closed
)

const (
	StatusDraftStr     string = "draft"
	StatusPublishedStr string = "publised"
	StatusClosedStr    string = "closed"
)

func NewStatus(status int) (Status, error) {
	if status != int(Draft) && status != int(Published) {
		return Status(0), xerrors.Errorf("status must be %d or %d", Draft, Published)
	}

	return Status(status), nil
}

func NewStatusForUpdate(status int) (Status, error) {
	if status < int(Draft) {
		return Status(0), xerrors.Errorf("status must be over %d", Draft)
	}

	if status > int(Closed) {
		return Status(0), xerrors.Errorf("status must be less than %d", Closed)
	}

	return Status(status), nil
}

func (s Status) Value() int {
	return int(s)
}

func (s Status) String() string {
	switch s {
	case Draft:
		return StatusDraftStr
	case Published:
		return StatusPublishedStr
	case Closed:
		return StatusClosedStr
	default:
		return "Undefined"
	}
}

func (s Status) Equals(s2 Status) bool {
	return s == s2
}
