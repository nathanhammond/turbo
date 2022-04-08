# Contributing

We're excited that you're interested in contributing to Turborepo! If you're coming from the web world, your computer might not be configured to work on Turborepo itself.

We'll give you a quick crash course on getting set up.

## Set up Go

1. Grab a package from the [Go Download Page](https://go.dev/dl/).
2. Install the Delve debugger: `go install github.com/go-delve/delve/cmd/dlv@latest`

## Set up VSCode

1. Install the [VSCode Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go).

## Set up your environment

1. `brew install --ignore-dependencies sponge pnpm yarn`

## Development Environment Setup

Dependencies

1.  On OSX: `brew install sponge`
2.  Run `pnpm install` at root

Building

- Building turbo CLI: In `cli` run `make turbo`
- Using turbo to build turbo CLI: `./turbow.sh`

Smoke Testing via examples:

1.  In `cli` run `make e2e`

## Debugging

1.  Install `go get dlv-dap`
2.  In VS Code Debugging tab, select `Basic Turbo Build` to start debugging the initial launch of `turbo` against the `build` target of the Basic Example.
