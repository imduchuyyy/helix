import { createPublicClient, type PublicClient, http, createWalletClient, type Account, keccak256, type Hex } from "viem";
import { Action } from "./base";
import * as chains from "viem/chains";
import { privateKeyToAccount } from "viem/accounts";
import type { ITokenBalance } from "../types/token";

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

  async fetchTokenBalances(): Promise<ITokenBalance[]> {
    return []
  }
}
