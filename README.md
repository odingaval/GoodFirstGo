# GoodFirstGo 

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

> A powerful CLI tool to help developers discover beginner-friendly open-source issues on GitHub

**GoodFirstGo** makes it easy to find your next open-source contribution by searching GitHub for issues labeled with `good-first-issue`, `help-wanted`, and other beginner-friendly tags. Perfect for developers looking to contribute to open source!

---

## ✨ Features

### Current Features (MVP)
-  **Smart Issue Search** - Find GitHub issues by language, labels, and activity
- **Advanced Filtering** - Filter by repository stars, forks, and issue age
- **Detailed Display** - View issue title, repository, link, and labels in a clean CLI format
- **Save Favorites** - Bookmark interesting issues locally for later reference
- **Fast & Lightweight** - Built with Go for optimal performance

### Planned Features
- **Web Dashboard** - Interactive web interface with advanced search and filters
- **Notifications** - Get alerts for new beginner-friendly issues in your favorite languages
- **GitHub Integration** - Authenticate to star, comment, and interact with issues directly
- **Analytics** - Discover which repositories have the most beginner-friendly issues

---

## Tech Stack

- **Language**: Go 1.21+
- **HTTP Client**: `net/http` (standard library)
- **JSON Parsing**: `encoding/json` (standard library)
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra) or [urfave/cli](https://github.com/urfave/cli)
- **GitHub API**: [go-github](https://github.com/google/go-github) (optional)

---

##  Installation

### Prerequisites
- Go 1.21 or higher
- GitHub Personal Access Token (for higher API rate limits)

### From Source
```bash
# Clone the repository
git clone https://github.com/odingaval/GoodFirstGo.git
cd GoodFirstGo

# Initialize Go module
go mod init github.com/odingaval/GoodFirstGo

# Install dependencies
go mod tidy

# Build the binary
go build -o goodfirstgo

# (Optional) Install globally
go install
```

---

## Usage

### Basic Search
```bash
# Search for Go issues with 'good-first-issue' label
goodfirstgo search --language go --label good-first-issue

# Search for Python issues with 'help-wanted' label
goodfirstgo search --language python --label help-wanted
```

### Advanced Filtering
```bash
# Filter by minimum stars
goodfirstgo search --language javascript --stars 100+

# Filter by issue age (recent issues only)
goodfirstgo search --language rust --age recent

# Combine multiple filters
goodfirstgo search --language go --label good-first-issue --stars 500+ --age recent
```

### Managing Favorites
```bash
# Save an issue to favorites
goodfirstgo save <issue-url>

# List saved issues
goodfirstgo list

# Remove from favorites
goodfirstgo remove <issue-id>
```

---

## 📅 Development Roadmap

| Week | Focus Area | Tasks |
|------|-----------|-------|
| **1-2** | Core CLI | Set up Go module, implement basic search, fetch GitHub issues, display results |
| **3** | Filters | Add filtering by stars/forks/date, refactor code structure |
| **4** | Favorites | Implement local storage (JSON/YAML), add save/list commands |
| **5** | Error Handling | Handle rate-limiting, HTTP errors, add logging and validation |
| **6** | Polish | Clean UI, comprehensive documentation, prepare for release |

---

##  Contributing

Contributions are welcome! This project itself is designed to help beginners contribute to open source, so we're especially friendly to first-time contributors.

### How to Contribute
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

---

## Debugging & Best Practices

This project emphasizes learning through debugging:
- Structured error handling using `errors.Is` and custom errors
- Comprehensive logging for HTTP requests and API responses
- Minimal reproducible debugging techniques
- Rate-limiting and retry logic for GitHub API

---

##  License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Contact

**Project Maintainer**: [@odingaval](https://github.com/odingaval)

**Project Link**: [https://github.com/odingaval/GoodFirstGo](https://github.com/odingaval/GoodFirstGo)

---

<div align="center">
Made with Val for the open-source community
</div>