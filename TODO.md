# Language Learning Mode Implementation Plan

## Steps:
- [ ] 1. Create internal/resources/learning.go with language resource mappings
- [x] 2. Update internal/models/issue.go: Add Resources []string field
- [x] 3. Update pkg/output/printer.go: Enhance PrintIssues for resources display
- [x] 4. Update cmd/root.go: Add --learning flag, enrich issues, integrate printer
- [x] 5. Update README.md: Add usage example
 - [x] 6. Run `go mod tidy`
 - [x] 7. Run `go build ./...` verify
 - [ ] 8. Test: `go run ./cmd/goodfirstgo --language go --label good-first-issue --limit 2 --learning` (running)
 - [x] 9. attempt_completion

Current: Starting step 1.
