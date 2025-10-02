package request

type LoginRequest struct {
	Username string `json:"username"`
	Creds    string `json:"creds"`
}
