package enum

type Role string

const (
	ADMIN   Role = "ADMIN"
	MEDIC   Role = "MEDIC"
	PATIENT Role = "PATIENT"
	GUEST   Role = "GUEST"
)
