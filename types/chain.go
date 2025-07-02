package types

type Chain struct {
	Name         string   `json:"name"`
	Rpcs         []string `json:"rpcs"`
	TokenListURL string   `json:"tokenListUrl,omitempty"`
}
