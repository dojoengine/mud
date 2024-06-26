import type { CommandModule } from "yargs";
import { loadConfig } from "@latticexyz/config/node";
import { World as WorldConfig } from "@latticexyz/world";
import { worldgen } from "@latticexyz/world/node";
import { getSrcDirectory } from "@latticexyz/common/foundry";
import path from "path";
import { rmSync } from "fs";
import { getExistingContracts } from "../utils/getExistingContracts";

type Options = {
  configPath?: string;
  clean?: boolean;
  srcDir?: string;
  config?: WorldConfig;
};

const commandModule: CommandModule<Options, Options> = {
  command: "worldgen",

  describe: "Autogenerate interfaces for Systems and World based on existing contracts and the config file",

  builder(yargs) {
    return yargs.options({
      configPath: { type: "string", desc: "Path to the MUD config file" },
      clean: {
        type: "boolean",
        desc: "Clear the worldgen directory before generating new interfaces (defaults to true)",
        default: true,
      },
    });
  },

  async handler(args) {
    await worldgenHandler(args);
    process.exit(0);
  },
};

export async function worldgenHandler(args: Options) {
  const srcDir = args.srcDir ?? (await getSrcDirectory());

  const existingContracts = getExistingContracts(srcDir);

  // Load the config
  const mudConfig = args.config ?? ((await loadConfig(args.configPath)) as WorldConfig);

  const outputBaseDirectory = path.join(srcDir, mudConfig.codegen.outputDirectory);

  // clear the worldgen directory
  if (args.clean) {
    rmSync(path.join(outputBaseDirectory, mudConfig.codegen.worldgenDirectory), { recursive: true, force: true });
  }

  // generate new interfaces
  await worldgen(mudConfig, existingContracts, outputBaseDirectory);
}

export default commandModule;
