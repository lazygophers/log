---
name: 🐛 Bug Report
about: Report a bug to help us improve
title: '[BUG] '
labels: bug, needs-triage
assignees: ''
---

## 🐛 Bug Description
<!-- A clear and concise description of what the bug is -->

## 🔄 Steps to Reproduce
<!-- Steps to reproduce the behavior -->

1. 
2. 
3. 
4. 

## ✅ Expected Behavior
<!-- A clear and concise description of what you expected to happen -->

## ❌ Actual Behavior
<!-- A clear and concise description of what actually happened -->

## 🖼️ Screenshots/Logs
<!-- If applicable, add screenshots or log outputs to help explain your problem -->

```
Paste logs here
```

## 🌍 Environment
<!-- Please complete the following information -->

**Operating System:**
- OS: [e.g. Ubuntu 20.04, macOS 12.0, Windows 11]
- Architecture: [e.g. amd64, arm64]

**Go Environment:**
- Go version: [e.g. go1.21.0]
- GOOS: [e.g. linux, darwin, windows]
- GOARCH: [e.g. amd64, arm64]

**Library Information:**
- LazygoPHers/log version: [e.g. v1.0.0 or commit hash]
- Build tags used: [e.g. debug, release, discard, or none]
- Installation method: [e.g. go get, go mod]

**Dependencies:**
- Zap version: [if using Zap integration]
- Other relevant dependencies: 

## 📁 Code Sample
<!-- Minimal code sample that reproduces the issue -->

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Your code here that demonstrates the issue
}
```

## 🔍 Additional Context
<!-- Add any other context about the problem here -->

### Build Configuration
<!-- If relevant, include your build configuration -->

```bash
go build -tags="your-tags-here"
```

### Performance Impact
<!-- If this is a performance-related bug -->

- [ ] High CPU usage
- [ ] High memory usage  
- [ ] Slow log writes
- [ ] Memory leaks
- [ ] Other: ___________

### Frequency
<!-- How often does this bug occur? -->

- [ ] Always
- [ ] Sometimes
- [ ] Rarely
- [ ] Unable to reproduce consistently

### Severity
<!-- What's the impact of this bug? -->

- [ ] 🔥 Critical - Blocks development/production
- [ ] 🚨 High - Significant impact on functionality
- [ ] ⚠️ Medium - Some impact on functionality  
- [ ] ℹ️ Low - Minor issue or cosmetic

### Workaround
<!-- Have you found any workaround for this issue? -->

## 🤝 Contributing
<!-- Are you willing to help fix this? -->

- [ ] I'm willing to submit a PR to fix this issue
- [ ] I need help understanding how to fix this
- [ ] I can provide more information if needed
- [ ] I can help with testing the fix

---

**Checklist before submitting:**
- [ ] I searched for similar issues and couldn't find any
- [ ] I provided all the requested information above
- [ ] I included a minimal code sample that reproduces the issue
- [ ] I tested with the latest version of the library