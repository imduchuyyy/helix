import prompts from "prompts";
import { hashMessage } from "viem";
import { generateMnemonic, privateKeyToAccount } from "viem/accounts";

const generateAccount = async () => {
  const { entropy } = await prompts({
    type: "password",
    name: "entropy",
    message: "Enter entropy:",
  });

  if (!entropy) {
    throw new Error("Entropy is required");
  }

  const entropyHash = hashMessage(entropy);

  return privateKeyToAccount(entropyHash)
}

const main = async () => {
  const account = await generateAccount();
  console.log("Address:", account.address);
};

main().catch((error) => {
  console.error("Error in main:", error);
  process.exit(1);
});
