# agmcp — Antigravity MCP Manager

`agmcp` is a zero-dependency command-line interface (CLI) tool designed for Windows to automate and manage Model Context Protocol (MCP) server installations for the Antigravity CLI. 

It acts as a configuration bridge, letting you easily register both local command-based and remote SSE-based MCP servers into the Antigravity configuration file (`%USERPROFILE%\.gemini\antigravity-cli\mcp_config.json`).

---

## Why `agmcp`?

When migrating from **Claude Code** or other Gemini CLI environments, configuring MCP servers can be friction-prone on Windows. Hand-editing JSON configurations frequently leads to syntax issues, incorrect paths, and broken formatting.

`agmcp` eliminates this friction by:
- Resolving home directory and path boundaries automatically.
- Safely reading, parsing, merging, and writing configuration files without discarding unrelated top-level settings.
- Supporting modern, remote SSE configurations using the `serverUrl` field instead of deprecated `url` keys.

---

## Installation & Compilation

Since `agmcp` is written in Go with zero external dependencies, you can compile it directly into a single self-contained executable on your system.

### Build from Source
Ensure you have [Go](https://go.dev/) installed, then run:

```bash
go build -o agmcp.exe main.go
```

---

## Usage

### 1. Add/Update a Local Command-based Server
To register a local command-based MCP server (e.g., the Postgres server running via `npx`):

```bash
.\agmcp.exe add my-test-server npx @modelcontextprotocol/server-postgres
```

This will inject or update the configuration block:
```json
"my-test-server": {
  "command": "npx",
  "args": [
    "@modelcontextprotocol/server-postgres"
  ]
}
```

### 2. Add/Update a Remote SSE Server
To register a remote SSE-based MCP server using the `--remote` flag:

```bash
.\agmcp.exe add --remote my-remote-server https://example.com/sse
```

This will inject or update the configuration block using the correct, non-deprecated key:
```json
"my-remote-server": {
  "serverUrl": "https://example.com/sse"
}
```

---

## Configuration Location
On Windows, `agmcp` modifies your configuration at:
```
%USERPROFILE%\.gemini\antigravity-cli\mcp_config.json
```
All edits preserve any existing configuration settings in that file.
