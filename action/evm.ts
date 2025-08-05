import { createPublicClient, type PublicClient, http, createWalletClient, type Account, keccak256, type Hex, erc20Abi, formatUnits } from "viem";
import { Action } from "./base";
import * as chains from "viem/chains";
import { privateKeyToAccount } from "viem/accounts";
import type { ITokenBalance } from "../types/token";
import { TOKEN_LIST_URL } from "../constants";
import axios from "axios";

export class EVM extends Action {
  private name: string;
  private client: PublicClient;
  private account: Account;

  constructor(entropy: string, name: string) {
    super();
    this.name = name;
    this.client = createPublicClient({
      chain: (chains as any)[name],
      transport: http()
    }) as PublicClient;
    const privateKey = this.generatePrivateKeyFromEntropy(entropy);
    this.account = privateKeyToAccount(privateKey);
  }

  private generatePrivateKeyFromEntropy(entropy: string): Hex {
    return keccak256(Buffer.from(entropy + "evm" + "helix-wallet"));
  }

  getAddress(): string {
    return this.account.address;
  }

  chainName(): string {
    return this.name;
  }

  formatBalance(balance: bigint, decimals: number): string {
    return parseFloat(formatUnits(balance, decimals)).toFixed(4);
  }

  async fetchTokenBalances(): Promise<ITokenBalance[]> {
    if (!TOKEN_LIST_URL[this.name]) {
      throw new Error(`Token list URL for chain ${this.name} is not defined`);
    }

    const tokenBalances: ITokenBalance[] = [];
    const tokenListUrl = TOKEN_LIST_URL[this.name];
    const res = await axios.get(tokenListUrl);
    const contracts = res.data.map((token: any) => ({
      address: token.address,
      abi: erc20Abi,
      functionName: "balanceOf",
      args: ["0x4FFF0f708c768a46050f9b96c46C265729D1a62f"],
    }));
    const balances = await this.client.multicall({
      contracts,
    });

    const nativeBalance = await this.client.getBalance({
      address: "0x4FFF0f708c768a46050f9b96c46C265729D1a62f",
    });

    if (nativeBalance > 0n) {
      tokenBalances.push({
        tokenAddress: "native",
        tokenName: (chains as any)[this.name].nativeCurrency.name,
        tokenSymbol: (chains as any)[this.name].nativeCurrency.symbol,
        tokenDecimals: 18,
        balance: nativeBalance,
        formattedBalance: this.formatBalance(nativeBalance, 18),
      });
    }

    for (let i = 0; i < balances.length; i++) {
      const balanceResult = balances[i];
      if (
        balanceResult && balanceResult.status === "success" &&
        BigInt(balanceResult.result as bigint) > 0n
      ) {
        const token = res.data[i];
        tokenBalances.push({
          tokenAddress: token.address,
          tokenName: token.name,
          tokenSymbol: token.symbol,
          tokenDecimals: token.decimals,
          balance: BigInt(balanceResult.result as bigint),
          formattedBalance: this.formatBalance(BigInt(balanceResult.result as bigint), token.decimals),
        });
      }
    }

    return tokenBalances;
  }
}
