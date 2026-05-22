#!/bin/bash
set -e

BIN_DIR="$HOME/.local/bin"
mkdir -p "$BIN_DIR"

echo "Downloading agmcp for Linux..."
curl -fsSL "https://github.com/simplychiragk/agmcp/releases/latest/download/agmcp" -o "$BIN_DIR/agmcp"
chmod +x "$BIN_DIR/agmcp"

if [[ ":$PATH:" != *":$BIN_DIR:"* ]]; then
    echo "Adding $BIN_DIR to PATH in shell profile..."
    
    if [ -f "$HOME/.zshrc" ]; then
        PROFILE="$HOME/.zshrc"
    elif [ -f "$HOME/.bashrc" ]; then
        PROFILE="$HOME/.bashrc"
    else
        PROFILE="$HOME/.profile"
    fi
    
    echo "" >> "$PROFILE"
    echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$PROFILE"
    echo "Successfully updated $PROFILE. Please run 'source $PROFILE' or start a new terminal session."
else
    echo "agmcp is already in your PATH."
fi

echo "Installation complete! Try running 'agmcp list' in a new terminal."
