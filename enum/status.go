package enum

type Status int

const (
	TemporaryMember Status = iota + 1
	RegularMember
	ApplicationMember
	RejectMember
	ClosedMember
)
