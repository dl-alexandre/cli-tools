# Publishing cli-tools to GitHub

This guide walks you through publishing cli-tools as a public Go module.

## Prerequisites

- GitHub account with username `dl-alexandre`
- Git installed locally
- Go 1.25+ installed

## Steps to Publish

### 1. Create GitHub Repository

1. Go to https://github.com/new
2. Repository name: `cli-tools`
3. Description: "Shared Go library for CLI applications"
4. Visibility: **Public**
5. Do NOT initialize with README, .gitignore, or License (we have those files)
6. Click **Create repository**

### 2. Initialize and Push Code

From this directory (`/Users/developer/Documents/GitHub/workspaces/CLI-Tools/cli-tools`):

```bash
# Initialize git repository
git init

# Add all files
git add .

# Create initial commit
git commit -m "Initial commit: cli-tools v0.0.0

- Version management with semantic versioning
- GitHub release update checking
- Output formatting (table/JSON)
- Configuration management
- Kong CLI framework helpers
- Comprehensive test suite
- CI/CD workflows"

# Add remote (replace with your actual GitHub URL)
git remote add origin https://github.com/dl-alexandre/cli-tools.git

# Push to GitHub
git push -u origin main
```

### 3. Create v0.0.0 Release

```bash
# Create and push tag
git tag v0.0.0
git push origin v0.0.0
```

This triggers the release workflow and makes the module available via Go's module proxy.

### 4. Verify Release

Check that the release was created:
- https://github.com/dl-alexandre/cli-tools/releases

### 5. Update Local-UniFi-CLI

In `/Users/developer/Documents/GitHub/workspaces/CLI-Tools/Tools/Local-UniFi-CLI`:

```bash
# Remove the replace directive
go mod edit -dropreplace github.com/dl-alexandre/cli-tools

# Get the published version
go get github.com/dl-alexandre/cli-tools@v0.0.0

# Tidy dependencies
go mod tidy

# Test build
go build ./cmd/unifi
```

### 6. Update Other CLIs

Repeat step 5 for each CLI you want to migrate:
- `cimis-cli`
- `MyMarketNews-CLI`
- `UniFi-Site-Manager-CLI`
- etc.

## Post-Publish Checklist

- [ ] GitHub repository created and public
- [ ] Code pushed to main branch
- [ ] v0.0.0 tag created and pushed
- [ ] Release automatically created by GitHub Actions
- [ ] Local-UniFi-CLI successfully uses published module
- [ ] Go module proxy shows the package: https://pkg.go.dev/github.com/dl-alexandre/cli-tools

## Next Releases

After initial testing with your CLIs:

```bash
# Make changes, update CHANGELOG.md
git add .
git commit -m "Add feature X"
git push origin main

# Tag new version
git tag v0.1.0
git push origin v0.1.0

# Update CLIs
go get github.com/dl-alexandre/cli-tools@v0.1.0
```

## Troubleshooting

### Module not found
Wait 5-10 minutes after pushing the tag. Go's module proxy needs time to index.

### Import errors
Ensure the import path matches exactly:
```go
import "github.com/dl-alexandre/cli-tools/version"
```

### Build failures
Check that all dependencies are compatible with Go 1.25

## Version Strategy

- **v0.0.x** - Bug fixes and testing
- **v0.1.0** - First stable API (after all CLIs work)
- **v1.0.0** - Production ready with API stability guarantees

Keep v0.x.x until all dl-alexandre CLIs successfully migrate and work correctly.
