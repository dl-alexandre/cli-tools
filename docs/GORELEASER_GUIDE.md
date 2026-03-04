# GoReleaser Configuration Guide

This document outlines the GoReleaser configuration requirements to ensure compatibility with the Homebrew tap automation workflow.

## Archive Naming Convention

The automation workflow expects release assets named in this exact format:

```
{formula}-{os}-{arch}.tar.gz
```

Where:
- `{formula}` = the tool name (e.g., `abc`, `ams`, `gdrv`)
- `{os}` = `darwin` (macOS) or `linux`
- `{arch}` = `arm64` or `amd64`

### Examples:
- `abc-darwin-arm64.tar.gz`
- `abc-darwin-amd64.tar.gz`
- `abc-linux-arm64.tar.gz`
- `abc-linux-amd64.tar.gz`

## Recommended GoReleaser Configuration

Add this to your `.goreleaser.yml`:

```yaml
builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
    goarch:
      - amd64
      - arm64
    # IMPORTANT: Binary name inside archive must match formula name
    binary: "{{ .ProjectName }}"
    # Optional: Add ldflags for version info
    ldflags:
      - -s -w -X main.version={{ .Version }}

archives:
  - format: tar.gz
    # This naming template matches the automation workflow expectations
    name_template: "{{ .ProjectName }}-{{ .Os }}-{{ .Arch }}"
    files:
      - LICENSE*
      - README*
      - completions/*
    # Ensure the binary is at the root of the archive
    wrap_in_directory: false

# Homebrew tap integration (optional - if you want goreleaser to update formulas)
# Note: Our workflow handles this automatically, so this section is optional
brews:
  - name: {{ .ProjectName }}
    tap:
      owner: dl-alexandre
      name: homebrew-tap
    # Skip if you prefer the Ruby-based automation workflow
    skip_upload: true
```

## Verification Checklist

Before publishing a release, verify:

1. **Archive names match the pattern**: Run `goreleaser build --snapshot` and check `dist/` directory
2. **Binary inside archive is correctly named**: Extract a test archive and verify the binary name matches the formula name
3. **All four platforms are built**: Check that all combinations of darwin/linux + arm64/amd64 are present

## Testing the Automation

After publishing a release:

1. **Check GitHub Releases**: Verify the 4 tar.gz files are present with correct names
2. **Trigger the workflow**: The workflow runs daily, or you can trigger it manually at:
   `https://github.com/dl-alexandre/CLI-Tools/actions/workflows/update-formulas.yml`
3. **Review the PR**: The workflow will create a PR that replaces `TO_BE_UPDATED_*` with actual SHA256 values
4. **Test the formula**: The PR will trigger the `test-formulas` job which installs each formula on macOS and Linux

## Troubleshooting

### Workflow fails with "Failed to download core platforms"

This means the workflow couldn't download one of the required archives. Check:
- Did GoReleaser upload the assets?
- Are the file names exactly matching the pattern `{formula}-{os}-{arch}.tar.gz`?
- Is the release published (not a draft)?

### SHA256 values not being updated

The regex expects this pattern in the formula:
```ruby
url ".../abc-darwin-arm64.tar.gz"
sha256 "TO_BE_UPDATED_DARWIN_ARM64"
```

If the formula uses raw binaries (without .tar.gz) or different formatting, the regex won't match.

### Formula installs but `version` command fails

Ensure your CLI returns exit code 0 for the version command. The test job will fail if `--version` returns non-zero.

## Migration Status

The following formulas have been standardized to use tar.gz archives:

| Formula | Status | Archive Format |
|---------|--------|----------------|
| abc.rb | ✅ Migrated | tar.gz |
| ams.rb | ✅ Migrated | tar.gz (was source) |
| ask.rb | ✅ Migrated | tar.gz (was source) |
| cimis.rb | ✅ Already binary | tar.gz |
| cli-template.rb | ✅ Migrated | tar.gz |
| gdrv.rb | ✅ Already binary | tar.gz |
| gpd.rb | ✅ Already binary | tar.gz |
| mpr.rb | ✅ Already binary | tar.gz |
| unifi.rb | ✅ Migrated | tar.gz (was raw binary) |
| usm.rb | ✅ Migrated | tar.gz (was raw binary) |

All formulas now use the standardized `bin.install "name"` pattern.
