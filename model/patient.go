package model

type Patient struct {
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Cellphone string `json:"cellphone"`
}
