# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a favicon hash calculator that computes Shodan-compatible MurmurHash3 values for favicon files. It includes both Go and Python implementations for generating hashes used in Shodan search queries.

## Architecture

- **main.go**: Go implementation with custom MurmurHash3 algorithm that matches Python's mmh3 library behavior
- **get-shodan-favicon-hash.py**: Python reference implementation using the mmh3 library
- **favicon-hash**: Compiled Go binary for the main program

The Go implementation includes a custom MurmurHash3 function that ensures compatibility with Python's mmh3.hash() output, including proper handling of Base64 encoding with newlines every 76 characters.

## Development Commands

### Build
```bash
go build -o favicon-hash main.go
```

### Run Go version
```bash
./favicon-hash <favicon_url>
./favicon-hash -hex <favicon_url>
```

### Run Python version
```bash
python get-shodan-favicon-hash.py
```

### Dependencies
```bash
go mod tidy
```

## Key Implementation Details

- The Go version implements MurmurHash3 manually to ensure exact compatibility with Python's mmh3 library
- Base64 encoding includes newlines every 76 characters to match Python's codecs.encode behavior
- Both implementations should produce identical hash values for the same favicon
- The hash output can be formatted as decimal (default) or hexadecimal (-hex flag)