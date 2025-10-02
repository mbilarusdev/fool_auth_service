package request

type RegisterRequest struct {
	Username string `json:"username"`
	Creds    string `json:"creds"`
}
