package view

type UserCreateView struct {
	Email    string `json:"email" `
	Name     string `json:"name"`
	Password string `json:"password"`
}
