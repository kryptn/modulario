package proto

type JsonLoginRequest struct {
	User string `json:"user"`
	Password string `json:"password"`
}

type JsonCreateRequest struct {
	User  string   `json:"user"`
	Links []string `json:"links"`
	DeciderType string `json:"decider_type"`
}
