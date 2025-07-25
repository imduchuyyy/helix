export interface ITokenBalance {
  tokenAddress: string;
  tokenName: string;
  tokenSymbol: string;
  tokenDecimals: number;
  balance: bigint;
}
