# GoodFirstGo

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

CLI tool to discover beginner-friendly GitHub issues.

## Current Features

- **Smart Search**: By language + labels (good-first-issue, help-wanted)
- **Filtering**: `--stars >100`, `--age recent/week/month`
- **Learning Mode**: `--learning` shows lang tutorials/resources
- **Config**: GITHUB_TOKEN for rate limits
- **Fast**: Stdlib HTTP, clean table output

## Install

```bash
git clone <repo>
cd GoodFirstGo
go mod tidy
go build -o goodfirstgo ./cmd/goodfirstgo
./goodfirstgo --help
```

## Usage

**Basic**:
```bash
./goodfirstgo --language go --label "good-first-issue" --limit 5
```

**Advanced Filtering**:
```bash
./goodfirstgo --language python --stars 100 --age week
```

**Learning Mode** (shows language-specific resources):
```bash
./goodfirstgo --language rust --learning
```

### Full Options
- `--limit`: Max issues to fetch (default 5)
- `--language`: Target programming language (default "go")
- `--label`: Issue label to search for (default "good-first-issue")
- `--stars`: Minimum repository stars (default 0)
- `--age`: Maximum issue age: recent/week/month
- `--learning`: Enable learning mode to show language resources
