package evm

// package actions

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"math/big"
// 	"net/http"
// 	"sync"

// 	"github.com/ethereum/go-ethereum"
// 	ethcommon "github.com/ethereum/go-ethereum/common"
// 	ethtypes "github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/forta-network/go-multicall"
// 	"github.com/imduchuyyy/helix-wallet/common"
// 	"github.com/imduchuyyy/helix-wallet/types"
// )

// type balanceOutput struct {
// 	Balance *big.Int
// }

// func (a *Action) fetchTokenList(tokenListRpc string) ([]types.Token, error) {
// 	response, err := http.Get(tokenListRpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var tokenList []types.Token
// 	err = json.Unmarshal(body, &tokenList)
// 	if err != nil {
// 		return nil, fmt.Errorf("error unmarshalling token list: %w", err)
// 	}

// 	return tokenList, nil
// }

// func (a *Action) fetchTokenBalances(address ethcommon.Address) ([]types.TokenWithBalance, error) {
// 	tokenList, err := a.fetchTokenList(a.chain.TokenListURL)
// 	if err != nil {
// 		return nil, err
// 	}

// 	caller, err := multicall.Dial(context.Background(), a.chain.Rpcs[0])
// 	if err != nil {
// 		return nil, fmt.Errorf("error connecting to multicall RPC: %w", err)
// 	}
// 	// Initialize an array to hold all contract calls
// 	var contractCalls []*multicall.Call

// 	for _, token := range tokenList {
// 		contract, err := multicall.NewContract(common.ERC20ABI, token.Address)
// 		if err != nil {
// 			return nil, fmt.Errorf("error creating contract for token %s: %w", token.Symbol, err)
// 		}

// 		call := contract.NewCall(
// 			new(balanceOutput),
// 			"balanceOf",
// 			address,
// 		).Name(token.Symbol)

// 		contractCalls = append(contractCalls, call)
// 	}

// 	// Split into batches of maximum 100 calls
// 	var filteredTokens []types.TokenWithBalance
// 	batchSize := 100

// 	var mu sync.Mutex
// 	var wg sync.WaitGroup
// 	var callErrors []error

// 	// Start a goroutine to fetch ETH balance in parallel with token balances
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		client, err := ethclient.Dial(a.chain.Rpcs[0])
// 		if err != nil {
// 			fmt.Println("failed to get ETH client:", err.Error())
// 			return
// 		}

// 		balance, err := client.BalanceAt(context.Background(), address, nil)
// 		if err != nil {
// 			fmt.Println("failed to fetch ETH balance:", err.Error())
// 			return
// 		}

// 		if balance.Cmp(big.NewInt(0)) > 0 {
// 			// If ETH balance is greater than 0, add it to the filtered tokens
// 			filteredTokens = append(filteredTokens, types.TokenWithBalance{
// 				Detail:  types.Token{Address: common.ZERO_ADDRESS.String(), Symbol: "ETH", Name: "Ether", Decimals: 18},
// 				Balance: balance,
// 			})
// 		}
// 	}()

// 	for i := 0; i < len(contractCalls); i += batchSize {
// 		wg.Add(1)
// 		go func(startIdx int) {
// 			defer wg.Done()

// 			end := min(startIdx+batchSize, len(contractCalls))

// 			batchCalls, err := caller.Call(nil, contractCalls[startIdx:end]...)
// 			if err != nil {
// 				mu.Lock()
// 				callErrors = append(callErrors, fmt.Errorf("error executing multicall batch %d: %w", startIdx/batchSize, err))
// 				mu.Unlock()
// 				return
// 			}

// 			mu.Lock()
// 			for i, call := range batchCalls {
// 				balance := call.Outputs.(*balanceOutput).Balance
// 				if balance.Cmp(big.NewInt(0)) > 0 {
// 					tokenIndex := startIdx + i
// 					newToken := types.TokenWithBalance{
// 						Detail:  tokenList[tokenIndex],
// 						Balance: balance,
// 					}
// 					filteredTokens = append(filteredTokens, newToken)
// 				}
// 			}
// 			mu.Unlock()
// 		}(i)
// 	}

// 	wg.Wait()

// 	if len(callErrors) > 0 {
// 		return nil, callErrors[0]
// 	}

// 	return filteredTokens, nil
// }

// func (a *Action) transferToken(tokenAddress ethcommon.Address, fromAddress ethcommon.Address, toAddress ethcommon.Address, amount *big.Int) error {
// 	client, err := ethclient.Dial(a.chain.Rpcs[0])
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to Ethereum client: %w", err)
// 	}

// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		return fmt.Errorf("failed to get nonce for address %s: %w", fromAddress.Hex(), err)
// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		return fmt.Errorf("failed to suggest gas price: %w", err)
// 	}

// 	transferFnSignature := []byte("transfer(address,uint256)")
// 	hash := crypto.Keccak256(transferFnSignature)
// 	methodID := hash[:4]

// 	paddedAddress := ethcommon.LeftPadBytes(toAddress.Bytes(), 32)
// 	paddedAmount := ethcommon.LeftPadBytes(amount.Bytes(), 32)

// 	var data []byte
// 	data = append(data, methodID...)
// 	data = append(data, paddedAddress...)
// 	data = append(data, paddedAmount...)

// 	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
// 		To:   &toAddress,
// 		Data: data,
// 	})

// 	if err != nil {
// 		return fmt.Errorf("failed to estimate gas: %w", err)
// 	}

// 	tx := ethtypes.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)
// 	chainID, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		return fmt.Errorf("failed to get chain ID: %w", err)
// 	}

// 	signedTx, err := a.keyring.SignTransaction(tx, chainID)
// 	if err != nil {
// 		return fmt.Errorf("failed to sign transaction: %w", err)
// 	}

// 	fmt.Print("Signed transaction: ", signedTx.Hash().Hex(), "\n")

// 	return nil
// }
