import { IAction } from "../types/action"
import { EVM } from "../evm";

interface Chains {
  [key: string]: (entropy: string) => IAction
}

export const CHAINS: Chains = {
  mainnet: (entropy: string) => {
    return new EVM(entropy, "mainnet", "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID", "https://tokens.coingecko.com/uniswap/all.json");
  }
}
