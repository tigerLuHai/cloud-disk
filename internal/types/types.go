// Code generated by goctl. DO NOT EDIT.
package types

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Result
}

type RegisterRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type RegisterResponse struct {
	Result
}
