package handler

import "github.com/imduchuyyy/helix-wallet/types"

type Handler struct {
	action types.Action
}

func New(action types.Action) *Handler {
	return &Handler{
		action: action,
	}
}
