# Changelog

All notable changes to this project are documented here. This project follows [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) and [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-11-26

### Added
- `retry` for context-aware operation retries with bounded exponential backoff, jitter, and `Retry-After` support.
- `appleauth` for ES256 JWT signing from PEM bytes or files, with early-refresh token caching.

### Changed
- Document the supported package set as `version`, `update`, `cache`, `retry`, and `appleauth`.
- Remove unused package implementations and their dependencies.
- Keep existing `version`, `update`, and `cache` public entry points available.

## [0.0.1] - 2024-03-30

### Added
- Initial development release for the CLI tools fleet.

[Unreleased]: https://github.com/dl-alexandre/cli-tools/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/dl-alexandre/cli-tools/compare/v0.0.1...v0.1.0
[0.0.1]: https://github.com/dl-alexandre/cli-tools/releases/tag/v0.0.1
