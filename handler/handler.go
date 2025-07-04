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

func (h *Handler) handleGetAddress(args []string) error {
	address, err := h.action.GetAddress()
	if err != nil {
		return err
	}
	println("Wallet address:", address)
	return nil
}

func (h *Handler) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "get_address",
			Description: "Get the wallet address",
			Handler:     h.handleGetAddress,
			Usage:       "get_address",
		},
	}

}
