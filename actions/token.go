package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"sync"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/forta-network/go-multicall"
	"github.com/imduchuyyy/helix-wallet/common"
	"github.com/imduchuyyy/helix-wallet/types"
)

type balanceOutput struct {
	Balance *big.Int
}

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
	err = json.Unmarshal(body, &tokenList)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling token list: %w", err)
	}

	return tokenList, nil
}

func (a *Action) fetchTokenBalances(tokenListRpc string, rpcUrl string, address ethcommon.Address) ([]types.TokenWithBalance, error) {
	tokenList, err := a.fetchTokenList(tokenListRpc)
	if err != nil {
		return nil, err
	}

	caller, err := multicall.Dial(context.Background(), rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("error connecting to multicall RPC: %w", err)
	}
	// Initialize an array to hold all contract calls
	var contractCalls []*multicall.Call

	for _, token := range tokenList {
		contract, err := multicall.NewContract(common.ERC20ABI, token.Address)
		if err != nil {
			return nil, fmt.Errorf("error creating contract for token %s: %w", token.Symbol, err)
		}

		call := contract.NewCall(
			new(balanceOutput),
			"balanceOf",
			address,
		).Name(token.Symbol)

		contractCalls = append(contractCalls, call)
	}

	// Split into batches of maximum 100 calls
	var filteredTokens []types.TokenWithBalance
	batchSize := 100

	var mu sync.Mutex
	var wg sync.WaitGroup
	var callErrors []error

	// Start a goroutine to fetch ETH balance in parallel with token balances
	wg.Add(1)
	go func() {
		defer wg.Done()
		client, err := ethclient.Dial(rpcUrl)
		if err != nil {
			fmt.Println("failed to get ETH client:", err.Error())
			return
		}

		balance, err := client.BalanceAt(context.Background(), address, nil)
		if err != nil {
			fmt.Println("failed to fetch ETH balance:", err.Error())
			return
		}

		if balance.Cmp(big.NewInt(0)) > 0 {
			// If ETH balance is greater than 0, add it to the filtered tokens
			filteredTokens = append(filteredTokens, types.TokenWithBalance{
				Detail:  types.Token{Address: common.ZERO_ADDRESS.String(), Symbol: "ETH", Name: "Ether", Decimals: 18},
				Balance: balance,
			})
		}
	}()

	for i := 0; i < len(contractCalls); i += batchSize {
		wg.Add(1)
		go func(startIdx int) {
			defer wg.Done()

			end := min(startIdx+batchSize, len(contractCalls))

			batchCalls, err := caller.Call(nil, contractCalls[startIdx:end]...)
			if err != nil {
				mu.Lock()
				callErrors = append(callErrors, fmt.Errorf("error executing multicall batch %d: %w", startIdx/batchSize, err))
				mu.Unlock()
				return
			}

			mu.Lock()
			for i, call := range batchCalls {
				balance := call.Outputs.(*balanceOutput).Balance
				if balance.Cmp(big.NewInt(0)) > 0 {
					tokenIndex := startIdx + i
					newToken := types.TokenWithBalance{
						Detail:  tokenList[tokenIndex],
						Balance: balance,
					}
					filteredTokens = append(filteredTokens, newToken)
				}
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if len(callErrors) > 0 {
		return nil, callErrors[0]
	}

	return filteredTokens, nil
}
