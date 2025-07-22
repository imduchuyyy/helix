import { Cli } from "./cli";
import { genChainInstance } from "./chain";
import prompts from "prompts";

const main = async () => {
  const chainDenote = process.env.CHAIN;

  if (!chainDenote) {
    throw new Error("CHAIN environment variable is not set");
  }

  const getChainActions = genChainInstance(chainDenote);
  if (!getChainActions) {
    throw new Error(`Chain ${chainDenote} is not supported`);
  }

  const { entropy } = await prompts({
    type: "password",
    name: "entropy",
    message: "Enter entropy:",
  })

  if (!entropy) {
    throw new Error("Entropy is required");
  }


  const action = getChainActions(entropy);
  const cli = new Cli();
  cli.setPrompt(`helix-${action.chainName()}> `);
  cli.registerCommands(action.commands)

  await cli.run();
}

main().catch((error) => {
  console.error("Error in main:", error);
  process.exit(1);
});
