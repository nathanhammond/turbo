export type BaseVersions = "berry" | "yarn" | "pnpm" | "npm";
export type CommandName = "yarn" | "pnpm" | "npm";

export type PackageManager = {
  name: BaseVersions;
  command: CommandName;
  version: string;
};

export const PACKAGE_MANAGERS: Record<BaseVersions, PackageManager> = {
  npm: {
    name: "npm",
    command: "npm",
    version: "latest",
  },
  pnpm: {
    name: "pnpm",
    command: "pnpm",
    version: "latest",
  },
  yarn: {
    name: "yarn",
    command: "yarn",
    version: "1.x",
  },
  berry: {
    name: "berry",
    command: "yarn",
    version: "stable",
  },
};
