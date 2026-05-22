<h1 align="center">⚡ agmcp ⚡</h1>

<p align="center">
  <strong>The Ultimate Zero-Dependency MCP Manager for the Antigravity CLI</strong>
</p>

<p align="center">
  <img src="agmcp.gif" alt="agmcp demo" width="680" style="border-radius: 8px; box-shadow: 0 8px 24px rgba(0,0,0,0.15);" />
</p>

<p align="center">
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-4CAF50?style=for-the-badge" alt="License" /></a>
  <img src="https://img.shields.io/badge/OS-Windows%20%7C%20Linux%20%7C%20macOS-2196F3?style=for-the-badge" alt="OS Support" />
  <a href="https://github.com/simplychiragk/agmcp/pulls"><img src="https://img.shields.io/badge/PRs-welcome-ff69b4?style=for-the-badge" alt="PRs Welcome" /></a>
</p>

---

## 📦 Installation

Install `agmcp` with a single command:

#### Linux / macOS
```bash
curl -fsSL https://raw.githubusercontent.com/simplychiragk/agmcp/main/install.sh | bash
```

#### Windows (PowerShell)
```powershell
irm https://raw.githubusercontent.com/simplychiragk/agmcp/main/install.ps1 | iex
```

#### Windows (CMD)
```cmd
curl -fsSL https://raw.githubusercontent.com/simplychiragk/agmcp/main/install.cmd -o install.cmd && install.cmd && del install.cmd
```

---

## 🪝 The Hook: Manual Configs are a Nightmare

Windows users migrating to the **Antigravity CLI** from Claude Code or the Gemini CLI often hit a wall when setting up MCP servers:

* 😣 **Syntax Nightmares:** Editing hidden, deeply nested `mcp_config.json` files by hand inevitably leads to broken commas and invalid JSON.
* 🔍 **Path Hunting:** Finding and resolving `%USERPROFILE%` correctly across shells and environments is error-prone.
* ⚠️ **Schema Drift:** Keeping track of when to use `command`/`args` versus migrating to modern SSE-based remote `serverUrl` options.

**Enter `agmcp`.** A single compiled, zero-dependency binary that safely manages your MCP configurations directly from the CLI. Run one command, and your MCP server is instantly registered, formatted, and ready for action.

---

## ⚡ Quick Start

Get up and running in seconds. No node modules, no dependencies, just raw performance.

### Install & Build from Source
```bash
# Clone the repository and build the binary
git clone https://github.com/simplychiragk/agmcp.git
cd agmcp
go build -o agmcp.exe main.go
```

### Adding an MCP Server (e.g., GitHub MCP)

#### 💻 PowerShell
```powershell
.\agmcp.exe add github-server npx -y @modelcontextprotocol/server-github
```

#### 🐧 Bash (Linux / macOS)
```bash
./agmcp add github-server npx -y @modelcontextprotocol/server-github
```

### Adding a Remote SSE Server
```bash
./agmcp add --remote remote-prod https://mcp.yourdomain.com/sse
```

---

## ✨ Features

* 🚀 **Zero Dependencies:** Single static binary with zero external runtime requirements.
* 📂 **Path Safety:** Automatic resolution of Windows profile paths using native OS bindings.
* 🛡️ **Preservative Formatting:** Safely parses, updates, and pretty-prints your `%USERPROFILE%\.gemini\antigravity-cli\mcp_config.json` without wiping other configuration settings.
* 🌐 **Remote Native:** Modern support for the latest remote SSE protocols using the standard `serverUrl` schema (no more deprecated `url` keys!).

---

## 🌟 Show Your Support

If `agmcp` saved you from configuration headaches, please consider **starring the repository**! ⭐

Got features you want to see? Found a bug? Open an issue on our [GitHub Issues](https://github.com/simplychiragk/agmcp/issues) tab! Contributions, ideas, and PRs are always welcome. Let's make MCP configuration completely painless.
