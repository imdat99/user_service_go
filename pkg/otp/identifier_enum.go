package otp

type IdentifierType int

const (
	Email IdentifierType = iota
	Phone
	Username
)

func (t IdentifierType) String() string {
	switch t {
	case Email:
		return "email"
	case Phone:
		return "phone"
	case Username:
		return "username"
	default:
		return "unknown"
	}
}
