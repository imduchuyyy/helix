export interface IAction {
  getAddress(): string;
  getPrivateKey(): string;
  chainName(): string;
}
