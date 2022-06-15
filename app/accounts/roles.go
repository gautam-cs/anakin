package accounts

type SUserRoleType string

//Export UserStatus types
const (
	SUserRoleTypeUnknown SUserRoleType = "unknown"
	SUserRoleTypeRoot    SUserRoleType = "root"
	SUserRoleTypeUser    SUserRoleType = "user"
)

func (role SUserRoleType) IUserRoleType() IUserRoleType {
	switch role {
	case SUserRoleTypeUnknown:
		return IUserRoleTypeUnknown
	case SUserRoleTypeRoot:
		return IUserRoleTypeRoot
	case SUserRoleTypeUser:
		return IUserRoleTypeRoot
	default:
		return IUserRoleTypeUnknown
	}
}

func (role SUserRoleType) String() string {
	return string(role)
}

type IUserRoleType int

//Export UserStatus types
const (
	IUserRoleTypeUnknown IUserRoleType = iota
	IUserRoleTypeRoot
	IUserRoleTypeSUser
)

func (role IUserRoleType) SUserRoleType() SUserRoleType {
	switch role {
	case IUserRoleTypeUnknown:
		return SUserRoleTypeUnknown
	case IUserRoleTypeRoot:
		return SUserRoleTypeRoot
	case IUserRoleTypeSUser:
		return SUserRoleTypeUser
	default:
		return SUserRoleTypeUnknown
	}
}

func (role IUserRoleType) Int() int {
	return int(role)
}
