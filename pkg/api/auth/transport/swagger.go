package transport

import (
	types "github.com/noartem/godi-example"
)

// Login request
// swagger:parameters login
type swaggLoginReq struct {
	// in:body
	Body credentials
}

// Login response
// swagger:response loginResp
type swaggLoginResp struct {
	// in:body
	Body struct {
		*types.AuthToken
	}
}

// Register request
// swagger:parameters register
type swaggRegisterReq struct {
	// in:body
	Body newUser
}

// Register response
// swagger:response registerResp
type swaggRegisterResp struct {
	// in:body
	Body struct {
		*types.User
	}
}

// Refresh request
// swagger:parameters refresh
type swaggRefreshReq struct {
	// in:body
	Body refreshReq
}

// Token refresh response
// swagger:response refreshResp
type swaggRefreshResp struct {
	// in:body
	Body struct {
		*types.AuthToken
	}
}
