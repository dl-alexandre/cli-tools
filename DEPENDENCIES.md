# CLI-Tools Dependencies Audit

This document tracks all external Go dependencies used across the CLI-Tools monorepo. Dependencies are audited before starring on GitHub to ensure we actually need them and are using their latest features.

## Legend
- ⭐ **Starred** - Repository starred on GitHub
- ✅ **Current** - Using latest version
- ⚠️ **Update Available** - Newer version exists
- 🔍 **Needs Review** - Need to verify if still needed
- 🗑️ **Remove** - Consider removing/unused

## Core Dependencies

### CLI Frameworks

| Repository | Used By | Current Version | Latest Version | Status | Notes |
|------------|---------|-----------------|----------------|--------|-------|
| [alecthomas/kong](https://github.com/alecthomas/kong) | 10 projects | v1.10.0-v1.14.0 | v1.14.0 | ✅ | Primary CLI framework |
| [spf13/cobra](https://github.com/spf13/cobra) | 4 projects | v1.8.0-v1.9.0 | v1.9.0 | ✅ | Alternative CLI (legacy projects) |
| [spf13/viper](https://github.com/spf13/viper) | 8 projects | v1.19.0 | v1.19.0 | ✅ | Configuration management |

### Security & Authentication

| Repository | Used By | Current Version | Latest Version | Status | Notes |
|------------|---------|-----------------|----------------|--------|-------|
| [zalando/go-keyring](https://github.com/zalando/go-keyring) | 3 projects | v0.2.5 | v0.2.5 | ✅ | Secure credential storage (X-CLI, etc.) |
| [99designs/keyring](https://github.com/99designs/keyring) | 1 project | v1.2.2 | v1.2.2 | 🔍 | Alternative keyring - consolidate with zalando? |
| [golang-jwt/jwt](https://github.com/golang-jwt/jwt) | 2 projects | v5.2.1-v5.3.1 | v5.3.1 | ✅ | JWT handling for App Store APIs |
| [refraction-networking/utls](https://github.com/refraction-networking/utls) | 1 project | v1.3.3 | v1.6.1 | ⚠️ | TLS fingerprinting - check if needed |

### HTTP Clients & Browser Automation

| Repository | Used By | Current Version | Latest Version | Status | Notes |
|------------|---------|-----------------|----------------|--------|-------|
| [go-resty/resty](https://github.com/go-resty/resty) | 6 projects | v2.17.2 | v2.17.2 | ✅ | Primary HTTP client |
| [chromedp/chromedp](https://github.com/chromedp/chromedp) | 1 project | v0.14.2 | v0.16.1 | ⚠️ | Browser automation for OAuth - needs update |
| [chromedp/cdproto](https://github.com/chromedp/cdproto) | 1 project | v0.0.0-20250803... | Latest | 🔍 | Auto-updates with chromedp |

### Output & Formatting

| Repository | Used By | Current Version | Latest Version | Status | Notes |
|------------|---------|-----------------|----------------|--------|-------|
| [rodaine/table](https://github.com/rodaine/table) | 7 projects | v1.3.0 | v1.3.1 | ⚠️ | Table formatting - minor update available |
| [mattn/go-isatty](https://github.com/mattn/go-isatty) | 7 projects | v0.0.20 | v0.0.20 | ✅ | Terminal detection |
| [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) | 2 projects | - | Latest | 🔍 | Alternative table lib - consolidate to rodaine? |
| [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) | 1 project | v1.1.0 | v1.1.0 | ✅ | Terminal styling (Grokipedia) |
| [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) | 0 projects | - | Latest | 🗑️ | Listed but unused? Verify |
| [jedib0t/go-pretty](https://github.com/jedib0t/go-pretty) | 1 project | v6.6.5 | v6.6.7 | ⚠️ | Pretty printing (UPS-CLI) |
| [schollz/progressbar](https://github.com/schollz/progressbar) | 2 projects | v3.18.0 | v3.18.0 | ✅ | Progress bars |

### Data Processing

| Repository | Used By | Current Version | Latest Version | Status | Notes |
|------------|---------|-----------------|----------------|--------|-------|
| [gocarina/gocsv](https://github.com/gocarina/gocsv) | 1 project | v0.0.0-20240520... | Latest | ✅ | CSV handling (MyMarketNews) |
| [xuri/excelize](https://github.com/xuri/excelize) | 1 project | v2.9.0 | v2.9.0 | ✅ | Excel files (MyMarketNews) |
| [google/uuid](https://github.com/google/uuid) | 2 projects | v1.6.0 | v1.6.0 | ✅ | UUID generation |
| [itchyny/gojq](https://github.com/itchyny/gojq) | 1 project | v0.12.17 | v0.12.17 | ✅ | JSON querying |
| [browserutils/kooky](https://github.com/browserutils/kooky) | 1 project | v0.2.9 | v0.2.9 | ✅ | Browser cookie parsing |

### Utilities

| Repository | Used By | Current Version | Latest Version | Status | Notes |
|------------|---------|-----------------|----------------|--------|-------|
| [sahilm/fuzzy](https://github.com/sahilm/fuzzy) | 1 project | v0.1.1 | v0.1.1 | ✅ | Fuzzy matching (Grokipedia) |
| [skip2/go-qrcode](https://github.com/skip2/go-qrcode) | 1 project | v0.0.0-20200617... | Latest | ✅ | QR code generation (UniFi) |
| [pkg/browser](https://github.com/pkg/browser) | 1 project | v0.0.0-20240102... | Latest | ✅ | Open browser (OAuth) |
| [jotaen/kong-completion](https://github.com/jotaen/kong-completion) | 1 project | v0.0.0-20250211... | Latest | ✅ | Shell completions for kong |

### Testing

| Repository | Used By | Current Version | Latest Version | Status | Notes |
|------------|---------|-----------------|----------------|--------|-------|
| [stretchr/testify](https://github.com/stretchr/testify) | Multiple | v1.9.0-v1.10.0 | v1.10.0 | ✅ | Test assertions |

## Indirect Dependencies (Auto-managed)

These are pulled in automatically by direct dependencies. No need to star unless we use them directly.

| Repository | Pulled In By | Notes |
|------------|--------------|-------|
| danieljoos/wincred | go-keyring | Windows credential store |
| godbus/dbus | go-keyring | Linux credential store |
| 99designs/go-keychain | keyring | macOS keychain (legacy) |
| fsnotify/fsnotify | viper | File watching |
| go-viper/mapstructure | viper | Map decoding |
| fatih/color | various | Terminal colors |

## Action Items

### Immediate
1. **Update chromedp** - v0.14.2 → v0.16.1 (breaking changes possible)
2. **Update rodaine/table** - v1.3.0 → v1.3.1 (safe update)
3. **Review utls** - Do we still need TLS fingerprinting?

### Consolidation Candidates
- Consider standardizing on one table library (rodaine vs olekukonko vs go-pretty)
- Consider standardizing on one keyring library (zalando vs 99designs)

### For Starring
After audit, star these high-impact repos:
1. alecthomas/kong (primary CLI framework)
2. zalando/go-keyring (security)
3. chromedp/chromedp (browser automation)
4. go-resty/resty (HTTP client)
5. spf13/viper (config management)
6. charmbracelet/* ecosystem (if we expand UI)

## Changelog

- **2026-03-30** - Initial dependency audit document created
- Next review: 2026-06-30 (quarterly)

---

**Note:** This document should be updated whenever:
- Adding new dependencies
- Updating existing dependencies
- Removing unused dependencies
- Starring repositories on GitHub
