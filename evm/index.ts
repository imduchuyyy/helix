import { createPublicClient, PublicClient, http, createWalletClient, Account, keccak256, Hex } from "viem";
import { Action } from "../action";
import * as chains from "viem/chains";
import { privateKeyToAccount } from "viem/accounts";

export class EVM extends Action {
  private name: string;
  private client: PublicClient;
  private account: Account;
  
  constructor(entropy: string, name: string) {
    super();
    this.name = name;
    this.client = createPublicClient({
      chain: chains[name],
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
}
