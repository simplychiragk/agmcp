package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// MCPServer represents the configuration of a single MCP server.
type MCPServer struct {
	Command   string   `json:"command,omitempty"`
	Args      []string `json:"args,omitempty"`
	ServerURL string   `json:"serverUrl,omitempty"`
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <command> [arguments]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  add    Add or update an MCP server configuration\n")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		flag.Usage()
	}

	cmd := os.Args[1]
	switch cmd {
	case "add":
		handleAdd()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", cmd)
		flag.Usage()
	}
}

func handleAdd() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	remoteOpt := addCmd.Bool("remote", false, "Register server as a remote SSE server using serverUrl")

	addCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s add [--remote] <server-name> <command/url> [args...]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

	// Parse flags following "add"
	if err := addCmd.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	args := addCmd.Args()
	if *remoteOpt {
		// Remote server URL addition
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: remote server registration requires both a server name and a URL.\n")
			addCmd.Usage()
		}
		serverName := args[0]
		serverURL := args[1]
		if len(args) > 2 {
			fmt.Fprintf(os.Stderr, "Warning: extra arguments ignored for remote server: %v\n", args[2:])
		}

		err := updateConfig(serverName, MCPServer{ServerURL: serverURL})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully added remote MCP server %q with URL %q\n", serverName, serverURL)
	} else {
		// Command-based server addition
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: command-based server registration requires both a server name and a command.\n")
			addCmd.Usage()
		}
		serverName := args[0]
		command := args[1]
		cmdArgs := args[2:]

		// Ensure we don't pass nil slice for args to produce empty JSON array [] instead of null
		if cmdArgs == nil {
			cmdArgs = []string{}
		}

		err := updateConfig(serverName, MCPServer{
			Command: command,
			Args:    cmdArgs,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully added MCP server %q with command %q and args %v\n", serverName, command, cmdArgs)
	}
}

func updateConfig(name string, newServer MCPServer) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("unable to resolve user home directory: %w", err)
	}

	configDir := filepath.Join(home, ".gemini", "antigravity-cli")
	configPath := filepath.Join(configDir, "mcp_config.json")

	// Read existing config into a generic map to preserve other top-level keys
	config := make(map[string]interface{})

	file, err := os.Open(configPath)
	if err == nil {
		dec := json.NewDecoder(file)
		if err := dec.Decode(&config); err != nil && err != io.EOF {
			file.Close()
			return fmt.Errorf("failed to parse existing config file: %w", err)
		}
		file.Close()
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to open config file: %w", err)
	}

	// Extract or initialize mcpServers
	var mcpServers map[string]interface{}
	if val, ok := config["mcpServers"]; ok {
		if m, ok := val.(map[string]interface{}); ok {
			mcpServers = m
		} else {
			mcpServers = make(map[string]interface{})
		}
	} else {
		mcpServers = make(map[string]interface{})
	}

	// Construct new server map
	serverMap := make(map[string]interface{})
	if newServer.ServerURL != "" {
		serverMap["serverUrl"] = newServer.ServerURL
	} else {
		serverMap["command"] = newServer.Command
		serverMap["args"] = newServer.Args
	}

	// Update the specific server
	mcpServers[name] = serverMap
	config["mcpServers"] = mcpServers

	// Create directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Write back configuration
	outFile, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file for writing: %w", err)
	}
	defer outFile.Close()

	enc := json.NewEncoder(outFile)
	enc.SetIndent("", "  ")
	if err := enc.Encode(config); err != nil {
		return fmt.Errorf("failed to encode configuration JSON: %w", err)
	}

	return nil
}
