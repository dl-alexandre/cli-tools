# Contributing to cli-tools

Thank you for your interest in contributing to cli-tools! This document provides guidelines and instructions for contributing.

## Development Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/dl-alexandre/cli-tools.git
   cd cli-tools
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run tests:**
   ```bash
   make test
   ```

## Making Changes

1. **Create a branch:**
   ```bash
   git checkout -b feature/my-feature
   ```

2. **Make your changes** following the coding standards below

3. **Add tests** for any new functionality

4. **Run the test suite:**
   ```bash
   make test
   make lint
   ```

5. **Commit your changes** with a descriptive message:
   ```bash
   git commit -m "Add feature: description of changes"
   ```

6. **Push to your fork** and submit a pull request

## Coding Standards

### Go Code Style

- Follow standard Go conventions (gofmt, golint)
- Use meaningful variable and function names
- Add comments for exported functions, types, and packages
- Keep functions focused and small
- Use table-driven tests where appropriate

### Testing Requirements

- All new code must have tests
- Maintain or improve test coverage
- Use `t.TempDir()` for temporary files
- Save and restore environment variables in tests
- Include both success and error test cases

### Documentation

- Update README.md if adding new features
- Add examples for new functionality
- Update CHANGELOG.md with your changes

## Pull Request Process

1. Update the README.md and/or MIGRATION.md with details of changes if applicable
2. Ensure all tests pass
3. Ensure code is properly formatted (`make fmt`)
4. Update CHANGELOG.md with your changes
5. The PR will be merged once it has been reviewed and approved

## Release Process

Releases are managed by maintainers:

1. Update CHANGELOG.md
2. Create a git tag: `git tag v0.x.x`
3. Push the tag: `git push origin v0.x.x`
4. GitHub Actions will create the release automatically

## Questions?

Feel free to open an issue for questions or discussion before submitting a PR.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
