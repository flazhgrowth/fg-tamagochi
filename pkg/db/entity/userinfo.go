package entity

type AccountInfo struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
