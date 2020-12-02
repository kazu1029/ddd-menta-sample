package mentor_recruitmentdm

import "golang.org/x/xerrors"

type Status int

const (
	Draft = iota + 1
	Published
	Closed
)

const (
	StatusDraftStr     string = "draft"
	StatusPublishedStr string = "publised"
	StatusClosedStr    string = "closed"
)

func NewStatus(status int) (Status, error) {
	if status < Draft {
		return Status(0), xerrors.Errorf("status must be over than %d", Draft)
	}

	if status > Closed {
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
