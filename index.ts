import { createPrompt } from "bun-promptx";
import { CHAINS } from "./constants";

const main = async () => {
  const chainDenote = process.env.CHAIN;

  if (!chainDenote) {
    throw new Error("CHAIN environment variable is not set");
  }

  const entropy = createPrompt("Enter entropy: ", {
    echoMode: 'password'
  })

  if (!entropy.value) {
    throw new Error("Entropy is required");
  }

  const getChainActions = CHAINS[chainDenote];
  if (!getChainActions) {
    throw new Error(`Chain ${chainDenote} is not supported`);
  }

  const action = getChainActions(entropy.value);
  console.log("Chain action created:", action);

}

main().catch((error) => {
  console.error("Error in main:", error);
  process.exit(1);
});
