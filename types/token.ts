export interface ITokenBalance {
  tokenAddress: string;
  tokenName: string;
  tokenSymbol: string;
  tokenDecimals: number;
  balance: bigint;
  formattedBalance?: string; // Optional formatted balance for display
}
