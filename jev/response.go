package jev

type Response struct {
	Answers map[string]Answer   `json:"answers"`
}
