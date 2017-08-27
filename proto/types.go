package proto

type JsonLoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type JsonCreateRequest struct {
	User        string     `json:"user"`
	Links       []JsonLink `json:"links"`
	DeciderType string     `json:"decider_type"`
}

type JsonLink struct {
	Url       string `json:"url"`
	Threshold uint   `json:"threshold"`
}
