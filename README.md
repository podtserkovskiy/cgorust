# CGO Rust Project with Buck2

Go HTTP server with Rust FFI integration, built with Buck2.

## Prerequisites

- **Buck2** - Build system
- **Rust** (edition 2024) - System Rust toolchain
- **Go** (1.25.5+) - System Go toolchain
- **cbindgen** - C header generation from Rust

## Setup

### Install Buck2

```bash
# macOS ARM64
brew install zstd
curl -L https://github.com/facebook/buck2/releases/download/2026-01-02/buck2-aarch64-apple-darwin.zst -o /tmp/buck2.zst
zstd -d /tmp/buck2.zst -o ~/.cargo/bin/buck2
chmod +x ~/.cargo/bin/buck2
buck2 --version
```

**Note:** Use pre-built binary. Building from source has compatibility issues.

### Install cbindgen

```bash
cargo install --force cbindgen
```

## Run

```bash
buck2 run //go:app
```

## Test

```bash
buck2 run //go:app & sleep 1 && curl http://localhost:8080/add && kill %1
```

Expected output: `5 + 7 = 12 (computed in Rust)`
