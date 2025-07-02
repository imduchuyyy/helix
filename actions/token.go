package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/imduchuyyy/helix-wallet/types"
)

func (a *Action) fetchTokenList(tokenListRpc string) ([]types.Token, error) {
	response, err := http.Get(tokenListRpc)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var tokenList []types.Token

	err = json.Unmarshal([]byte(string(body)), &tokenList)
	if err != nil {
		return nil, err
	}

	fmt.Print("Fetched token list: ", tokenList)

	return tokenList, nil
}
