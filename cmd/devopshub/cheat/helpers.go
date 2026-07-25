package cheat

// This file is intentionally minimal — it exists to document the package's
// internal conventions. Helpers that operate across cheatsheet subcommands
// are added here as the package grows.
//
// JSON mode detection: use printer.JSONMode() — the output.Printer exposes
// this accessor cleanly. No workaround is needed.
