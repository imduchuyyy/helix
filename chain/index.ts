import * as chains from "viem/chains";
import { Action } from "../action";
import { EVM } from "../evm";

export const genChainInstance = (chainName: string): (entropy: string) => Action => {
  // TODO: support more chain type
  if (!(chainName in chains)) {
    throw new Error(`Chain ${chainName} is not supported`);
  }

  return (entropy: string): Action => {
    return new EVM(entropy, chainName)
  }
}
