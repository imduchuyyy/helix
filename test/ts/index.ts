import { keccak256 } from "viem";
import { privateKeyToAccount } from "viem/accounts";

function generateAddress(entropy: string) {
  const seed = keccak256(Buffer.from(entropy + "evm" + "helix-wallet"));

  const account = privateKeyToAccount(seed);

  console.log("Generated Address:", account.address);
}

function main() {
  const entropy = "huy"; // Test entropy
  generateAddress(entropy);
}

main()