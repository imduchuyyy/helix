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
    }
  ]

  abstract getAddress(): string;
  abstract chainName(): string;
}
