package sdk

type Json map[string]any
type Form map[string]string

type AuthResp[T any] struct {
	State   int    `json:"state"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Error   string `json:"error"`
	Errno   int    `json:"errno"`
}
type FsResp[T any] struct {
	State   bool   `json:"state"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
