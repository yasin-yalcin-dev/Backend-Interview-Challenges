# Testing Scripts

This directory contains utility scripts for testing the Backend Interview Challenges project.

## Available Scripts

### `run_all_tests.sh`

This script runs all Go tests in the project across all problem categories.

**Features:**
- Automatically finds and runs tests for all problem categories
- Provides colored output for better readability
- Reports test failures with a summary at the end
- Returns non-zero exit code if any tests fail

**Usage:**
```bash
./scripts/run_all_tests.sh
```

### `test_cli.sh`

This script tests the CLI functionality by executing various commands.

**Features:**
- Builds the CLI application
- Tests commands for all problem categories
- Provides colored output for better readability
- Reports command failures with a summary at the end
- Cleans up build artifacts after testing
- Returns non-zero exit code if any commands fail

**Usage:**
```bash
./scripts/test_cli.sh
```

## Making Scripts Executable

If the scripts are not executable, you can make them executable with:

```bash
chmod +x scripts/run_all_tests.sh
chmod +x scripts/test_cli.sh
```

## Adding New Tests

When adding new problems to the project, you should:

1. For unit tests: Add `*_test.go` files to the problem directory
2. For CLI tests: Add new test commands to `test_cli.sh`

## Troubleshooting

If you encounter issues with the scripts:

- Make sure you run the scripts from the project root directory
- Ensure Go is properly installed and in your PATH
- Check if the project structure matches what the scripts expect
