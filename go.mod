module github.com/lazygophers/log

go 1.26.2

require (
	github.com/lazygophers/log/constant v0.0.0
	github.com/lazygophers/log/hooks v0.0.0
	github.com/petermattis/goid v0.0.0-20260113132338-7c7de50cc741
	go.uber.org/zap v1.27.1
)

require go.uber.org/multierr v1.10.0 // indirect

// Local development - use replace for subpackages
// Remove these lines after pushing tags to remote
replace github.com/lazygophers/log/constant => ./constant

replace github.com/lazygophers/log/hooks => ./hooks
