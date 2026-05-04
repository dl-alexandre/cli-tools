#!/bin/bash
# Publish cli-tools to GitHub
# Run this script to push cli-tools v0.0.0 to GitHub

set -e

CLI_TOOLS_DIR="/Users/developer/Documents/GitHub/workspaces/CLI-Tools/cli-tools"
GIT_DIR="$CLI_TOOLS_DIR/.git"

echo "=== Publishing cli-tools v0.0.0 to GitHub ==="
echo ""

# Check if GitHub CLI is installed (optional but helpful)
if command -v gh &> /dev/null; then
    echo "✓ GitHub CLI (gh) is installed"
    
    # Check if authenticated
    if gh auth status &> /dev/null; then
        echo "✓ GitHub CLI is authenticated"
        
        # Check if repo exists
        if gh repo view dl-alexandre/cli-tools &> /dev/null 2>&1; then
            echo "✓ Repository dl-alexandre/cli-tools exists"
        else
            echo "Creating GitHub repository..."
            gh repo create dl-alexandre/cli-tools --public --description "Shared Go library for CLI applications" || true
        fi
    else
        echo "⚠ GitHub CLI not authenticated. Run: gh auth login"
        echo "  Or manually create the repo at: https://github.com/new"
    fi
else
    echo "ℹ GitHub CLI not installed. You can install it with: brew install gh"
    echo "  Or manually create the repo at: https://github.com/new"
fi

echo ""
echo "=== Git Status ==="
GIT_DIR="$GIT_DIR" GIT_WORK_TREE="$CLI_TOOLS_DIR" git status

echo ""
echo "=== Pushing to GitHub ==="
echo "Running: git push -u origin main"
GIT_DIR="$GIT_DIR" GIT_WORK_TREE="$CLI_TOOLS_DIR" git push -u origin main || {
    echo ""
    echo "⚠ Push failed. This is expected if:"
    echo "  1. The GitHub repository doesn't exist yet"
    echo "  2. You're not authenticated with GitHub"
    echo ""
    echo "To fix:"
    echo "  1. Create the repo at: https://github.com/new"
    echo "     - Name: cli-tools"
    echo "     - Visibility: Public"
    echo "     - Don't initialize with README (we have one)"
    echo "  2. Run this script again, or manually run:"
    echo "     cd /Users/developer/Documents/GitHub/workspaces/CLI-Tools/cli-tools"
    echo "     git push -u origin main"
    exit 1
}

echo ""
echo "=== Pushing v0.0.0 Tag ==="
echo "Running: git push origin v0.0.0"
GIT_DIR="$GIT_DIR" GIT_WORK_TREE="$CLI_TOOLS_DIR" git push origin v0.0.0 || {
    echo "⚠ Tag push failed. You may need to push manually:"
    echo "     git push origin v0.0.0"
    exit 1
}

echo ""
echo "=== ✅ Successfully Published! ==="
echo ""
echo "Repository: https://github.com/dl-alexandre/cli-tools"
echo "Release:    https://github.com/dl-alexandre/cli-tools/releases/tag/v0.0.0"
echo "Go Module:  go get github.com/dl-alexandre/cli-tools@v0.0.0"
echo ""
echo "Next steps:"
echo "  1. Wait 5-10 minutes for Go's module proxy to index the package"
echo "  2. Update Local-UniFi-CLI to use the published version:"
echo "     cd /Users/developer/Documents/GitHub/workspaces/CLI-Tools/Tools/Local-UniFi-CLI"
echo "     go mod edit -dropreplace github.com/dl-alexandre/cli-tools"
echo "     go get github.com/dl-alexandre/cli-tools@v0.0.0"
echo "     go mod tidy"
echo "     go build ./cmd/unifi"
