import { formatUnits } from "viem";
import type { ITokenBalance } from "../types/token";

export abstract class Action {
  public commands = [
    {
      name: "address",
      description: "Get the address of the account",
      handler: async (_: string[]) => {
        const address = this.getAddress();
        if (!address) {
          throw new Error("Address not found. Ensure the account is initialized.");
        }

        console.log(`Address: ${address}`);
      },
      usage: "address"
    },
    {
      name: "chainName",
      description: "Get the name of the blockchain",
      handler: async () => {
        const name = this.chainName();
        if (!name) {
          throw new Error("Chain name not found. Ensure the account is initialized.");
        }

        console.log(`Chain Name: ${name}`);
      },
      usage: "chainName"
    },
    {
      name: "balance",
      description: "Get balance of the account",
      handler: async () => {
        const balances = await this.fetchTokenBalances();
        if (!balances || balances.length === 0) {
          console.log("No token balances found.");
          return;
        }

        console.log("Token Balances:");
        for (const balance of balances) {
          console.log(`- ${balance.tokenName} (${balance.tokenSymbol}): ${balance.formattedBalance}`);
        }
      },
      usage: "chainName"
    }
  ]

  abstract fetchTokenBalances(): Promise<ITokenBalance[]>;
  abstract getAddress(): string;
  abstract chainName(): string;
}
