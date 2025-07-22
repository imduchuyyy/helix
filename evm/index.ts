import { createPublicClient, PublicClient, http } from "viem";
import { IAction } from "../types/action";
import * as chains from "viem/chains";

export class EVM implements IAction {
  private entropy: string;
  private name: string;
  private rpc: string;
  private tokenListRpc: string;
  private client: PublicClient;
  
  constructor(entropy: string, name: string, rpc: string, tokenListRpc: string) {
    this.entropy = entropy;
    this.name = name;
    this.rpc = rpc;
    this.tokenListRpc = tokenListRpc;
    this.client = createPublicClient({
      chain: chains[name],
      transport: http()
    }) as PublicClient;
  }

  getAddress(): string {
    // Placeholder for address generation logic
    // This should return the address derived from the entropy
    return "0xYourGeneratedAddress"; // Replace with actual address generation logic
  }

  getPrivateKey(): string {
    // Placeholder for private key generation logic
    // This should return the private key derived from the entropy
    return "0xYourGeneratedPrivateKey"; // Replace with actual private key generation logic
  }

  chainName(): string {
    return this.name;
  }

}
