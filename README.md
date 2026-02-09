Project Plan: Open-Source Contribution Helper (Go)
Goal

Create a CLI tool (and optionally a web interface later) that helps developers find beginner-friendly issues in GitHub repositories.

Core Features (MVP)

Search GitHub issues by criteria

Language (Go, Python, etc.)

Labels (good-first-issue, help-wanted)

Activity (recently updated)

Display issue details in CLI

Title, repository, link, labels

Optional: brief description

Filter issues

By repository stars or forks

By issue age (recent vs old)

Optional: Save favorite issues

Local JSON or YAML file

Track issues you want to work on

Advanced / Future Features

Web dashboard with search and filters

Notifications (e.g., new beginner-friendly issues in chosen languages)

Integration with GitHub authentication (to allow starring, commenting)

Analytics: which repos have most beginner-friendly issues

Tech Stack / Libraries

Go standard library for HTTP requests: net/http

JSON parsing: encoding/json

CLI framework: cobra or urfave/cli

Optional: go-github (official GitHub API client for Go)

Step-by-Step Roadmap (6 Weeks)
Week 1 – 2: Core CLI

Set up a Go module (go mod init)

Implement a simple CLI that takes language and label as input

Fetch GitHub issues using GitHub API

Display issue title, repo, and URL in terminal

Week 3: Filters

Add filtering by stars, forks, and date

Refactor code into reusable functions and structs

Week 4: Local Favorites

Add ability to save favorite issues in JSON/YAML

Add CLI command to view saved issues

Week 5: Error Handling & Logging

Handle rate-limiting, HTTP errors, invalid inputs

Add logs for debugging

Week 6: Polish & Documentation

Clean CLI interface

Add README with examples

Optional: publish as a small Go binary on GitHub

Debugging & Learning Focus

Use logging to track HTTP requests and responses

Add structured error handling (errors.Is, custom errors)

Practice minimal reproducible debugging: isolate any issue in fetching or parsing API responses