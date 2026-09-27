package jev

type Request struct {
	Model     string              `json:"model,omitempty"`
	State     string              `json:"state"`
	Questions map[string]Question `json:"questions"`
}
