package keyring

func (k *Keyring) handleGetAddress(args []string) (string, error) {
	address, err := k.GetEVMAddress()
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}
