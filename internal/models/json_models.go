package models

type TodoRequest struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type SignUpRequest struct {
	UserName string `json:"user_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Response struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Token   *Token           `json:"token,omitempty"`
	Profile *ProfileResponse `json:"profile,omitempty"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ProfileResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"user_name"`
	Login string `json:"login"`
	Role  string `json:"role"`
}

type ChangePassRequest struct {
	Login   string `json:"login"`
	OldPass string `json:"old_password"`
	NewPass string `json:"new_password"`
}
