# CPW - Copy on Write

A CLI tool that watches for changes in files or directories and copies the changed files to a target destination.

## Features

- Watch a single file or directory for changes
- Automatically copy changed files to the target destination
- Preserve directory structure when copying from a directory to another
- Initial copy of all files when watching directories
- Auto-create destination directories as needed
- Built-in versioning support

## Installation

### From Source
```bash
# Clone the repository
git clone https://github.com/mxvsh/cpw.git
cd cpw

# Build the binary
make build

# Or build with a specific version
make build VERSION=v1.0.0

# Optionally move to a directory in your PATH
cp cpw /usr/local/bin/
```

### From Releases
You can download pre-built binaries for ARM64 and ARM architectures from the [GitHub Releases](https://github.com/mxvsh/cpw/releases) page.

Available platforms:
- Linux ARM64 (aarch64)
- Linux ARM (armv7)
- macOS ARM64 (Apple Silicon)

## Usage

```bash
# Basic usage
cpw <source> <destination>

# Check version
cpw -version
```

### Examples

Watch a single file and copy it when changed:
```bash
cpw /path/to/source/file.txt /path/to/destination/file.txt
```

Watch a directory and copy any changed files with the same directory structure:
```bash
cpw /path/to/source/dir /path/to/destination/dir
```

## How It Works

1. The program starts by checking if both source and destination locations exist
2. If watching a directory, it initially copies all files to the destination
3. The program then monitors the source for file changes (write or create events)
4. When a change is detected, the file is copied to the destination, preserving directory structure
5. If new directories are created, they are automatically added to the watch list

## Versioning

CPW includes built-in versioning that displays:
- Version number (from Git tag)
- Commit SHA
- Build date

This information is automatically included in builds when:
- Building with `make build VERSION=v1.0.0`
- Creating a release with GitHub Actions

You can view the version information by running:
```bash
cpw -version
```

## CI/CD

This project uses GitHub Actions to automatically build and release binaries for different architectures:
- When a tag is pushed with the format `v*` (e.g., `v1.0.0`), GitHub Actions will:
  - Build binaries for ARM64 and ARM architectures
  - Include version information in the builds
  - Create a GitHub release with these binaries
  - Generate release notes

## Dependencies

- github.com/fsnotify/fsnotify - File system notifications for Go 