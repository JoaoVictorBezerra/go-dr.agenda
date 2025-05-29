package enum

type Role string

const (
	ADMIN   Role = "ADMIN"
	MANAGER Role = "MANAGER"
	MEDIC   Role = "MEDIC"
	PATIENT Role = "PATIENT"
	GUEST   Role = "GUEST"
)
