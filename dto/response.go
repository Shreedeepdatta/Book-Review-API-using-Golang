package dto

type AuthResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsAuthor bool   `json:"is_author"`
	Token    string `json:"token,omitempty"`
	Message  string `json:"message"`
}
