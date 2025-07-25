import { Action } from "../action/base";
import { ICommand } from "../types";
import prompts from "prompts";

export class Cli {
  private prompt: string;
  private welcomeMsg: string;
  private commands: Map<string, ICommand>;

  constructor() {
    this.prompt = "> ";
    this.welcomeMsg = "Helix-wallet CLI - \nType 'help' to see available commands.";
    this.commands = new Map();

    this.registerCommands(this.Commands());
  }

  setPrompt(prompt: string) {
    this.prompt = prompt;
  }

  setWelcomeMessage(msg: string) {
    this.welcomeMsg = msg;
  }

  private registerCommand(cmd: ICommand) {
    if (this.commands.has(cmd.name)) {
      console.warn(`Warning: Command '${cmd.name}' is already registered and will be overwritten.`);
    }
    this.commands.set(cmd.name, cmd);
  }

  registerCommands(cmds: ICommand[]) {
    for (const cmd of cmds) {
      this.registerCommand(cmd);
    }
  }

  public async run() {
    console.log(this.welcomeMsg);

    while (true) {
      const { value } = await prompts({
        type: "text",
        name: "value",
        message: this.prompt,
      });
      if (!value) continue;

      const args = value.split(/\s+/);
      const cmdName = args[0];

      const cmd = this.commands.get(cmdName);
      if (cmd) {
        try {
          await cmd.handler(args.slice(1));
        } catch (err) {
          console.error("Error:", err);
        }
      } else {
        console.error(`Unknown command: ${cmdName}. Type 'help' to see available commands.`);
      }
    }
  }

  // Example built-in command definitions (you can expand this)
  private Commands(): ICommand[] {
    return [
      {
        name: "help",
        description: "Show all available commands",
        handler: async (args: string[]) => this.helpHandler(args),
      },
      {
        name: "exit",
        description: "Exits the application",
        handler: async (_: string[]) => {
          console.log("Exiting the application...");
          process.exit(0);
        },
      },
    ];
  }

  private async helpHandler(args: string[]): Promise<void> {
    if (args.length > 0) {
      const cmd = this.commands.get(args[0]);
      if (cmd) {
        console.log(`${cmd.name}: ${cmd.description}`);
        if (cmd.usage) {
          console.log(`Usage: ${cmd.usage}`);
        }
      } else {
        console.error(`Unknown command: ${args[0]}`);
      }
      return;
    }

    console.log("Available commands:");
    for (const [name, cmd] of this.commands.entries()) {
      console.log(`  ${name.padEnd(12)} - ${cmd.description}`);
    }
  }
}
