export interface ICommand {
  name: string;
  description: string;
  handler: (args: string[]) => Promise<void>;
  usage?: string;
}
