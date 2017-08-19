package proto

type JsonCreateRequest struct {
	User  string   `json:"user"`
	Links []string `json:"links"`
}
