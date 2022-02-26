# Multi-level demo for github.com/google/subcommands

This tutorial contains one branch per topic.


## Usage

Checkout each branch in turn to learn one level at a time.

- Ensure you have already installed:
  - A working Go SDK
  - The [`staticcheck`](https://staticcheck.io) linting tool
  - The `make` command
- Run `make` to see the new features in that branch


## Contents

1. simple subcommands
   1. basic usage with builtin commands
   2. creating commands with the procedural API
   3. passing non-CLI arguments
   4. command groups ← _you are here_
   5. adding command flags
   6. marking flags as important
2. reusing command code
3. commanders
   1. procedural vs object API
   2. creating commands with custom commanders
   3. creating a testable command structure
4. creating nested commands
5. beyond `NewCommander`
   1. controlling output
   2. introspecting commanders
   3. replacing builtins
