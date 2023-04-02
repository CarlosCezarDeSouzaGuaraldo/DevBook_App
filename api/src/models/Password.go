package models

// Password is the request update password format
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
