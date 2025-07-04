package evm

type EVMAction struct {
	entropy string
}

func New(entropy string) EVMAction {
	return EVMAction{
		entropy: entropy,
	}
}

// func (a *Action) handleTransfer(args []string) (string, error) {
// 	if len(args) < 3 {
// 		return "", fmt.Errorf("usage: transfer [amount] [token] [address]")
// 	}
// 	// amountStr := args[0]
// 	// tokenSymbolOrTokenAddress := args[1]
// 	// toAddress := args[2]
// 	address, err := a.chain.Keyring.GetAddress()
// 	if err != nil {
// 		return "", fmt.Errorf("failed to get EVM address: %w", err)
// 	}

// 	a.transferToken(common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d"), address, common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d"), big.NewInt(1000000000000000))
// 	return "Transfer token", nil
// }

// func (a *Action) handleBalance(args []string) (string, error) {
// 	address, err := a.keyring.GetAddress()
// 	if err != nil {
// 		return "", fmt.Errorf("failed to get EVM address: %w", err)
// 	}

// 	tokenWithBalance, err := a.fetchTokenBalances(address)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to fetch token list: %w", err)
// 	}

// 	for _, token := range tokenWithBalance {
// 		decimalBalance := new(big.Float).Quo(
// 			new(big.Float).SetInt(token.Balance),
// 			new(big.Float).SetFloat64(math.Pow10(int(token.Detail.Decimals))),
// 		)
// 		fmt.Printf("Token: %s, Balance: %.6f\n", token.Detail.Symbol, decimalBalance)
// 	}

// 	return "", nil
// }
