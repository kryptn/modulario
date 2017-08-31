package proto

type responseBase struct {
	Result string `json:"result"`
	Err    string `json:"err,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Result string `json:"result"`
	Err    string `json:"err,omitempty"`
}

type LogoutRequest struct {
}

type LogoutResponse struct {
	responseBase
}

type RegisterRequest struct {
	LoginRequest
}

type RegisterResponse struct {
	responseBase
}

type VisitPostRequest struct {
	Key string `json:"key"`
}

type VisitPostResponse struct {
	responseBase
	Link string `json:"link"`
}

type ViewPostRequest struct {
	Key string `json:"key"`
}

type ViewPostResponse struct {
	responseBase
	Links []Link `json:"links"`
	Key   string `json:"key"`
}

type Link struct {
	Url      string `json:"url"`
	Accesses uint   `json:"accesses"`
}

type CreatePostRequest struct {
	User        string        `json:"user"`
	Links       []LinkRequest `json:"links"`
	DeciderType string        `json:"decider_type"`
}

type LinkRequest struct {
	Url       string `json:"url"`
	Threshold uint   `json:"threshold"`
}

type CreatePostResponse struct {
	responseBase
	Key string `json:"key"`
}

type DeletePostRequest struct {
	Key string `json:"key"`
}

type DeletePostResponse struct {
	responseBase
}
