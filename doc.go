// Package log provides a high-performance structured logging library for Go.
//
// This package is built on top of uber-go/zap and provides an easy-to-use
// interface with excellent performance characteristics.
//
// # Main Types
//
// Logger - The core logging type with methods for all log levels
// LoggerWithCtx - A context-aware logger that supports context.Context
// HourlyRotator - Automatic log file rotation by hour and size
// AsyncWriter - Asynchronous buffered writer for high-throughput scenarios
//
// # Log Levels
//
// The package supports standard log levels (from lowest to highest):
//   - TraceLevel
//   - DebugLevel
//   - InfoLevel
//   - WarnLevel
//   - ErrorLevel
//   - PanicLevel (logs then panics)
//   - FatalLevel (logs then os.Exit(1))
//
// # Build Tags
//
// The package supports conditional compilation for different environments:
//   - debug: Enables all logging (default)
//   - release: Production mode with optimized performance
//   - discard: Disables all logging output
//   - canary: Similar to debug mode for canary deployments
//
// # Thread Safety
//
// All logger types are safe for concurrent use by multiple goroutines.
// The AsyncWriter uses internal buffering and goroutines for optimal
// performance in high-concurrency scenarios.
//
// # Basic Usage
//
//	package main
//
//	import "github.com/lazygophers/log"
//
//	func main() {
//	    // Use global logger
//	    log.Info("application started")
//	    log.Infof("listening on port %d", 8080)
//
//	    // Create custom logger
//	    logger := log.NewLogger()
//	    logger.SetLevel(log.DebugLevel)
//	    logger.Debug("debug message")
//	}
//
// # Context-Aware Logging
//
//	ctx := context.Background()
//	logger := log.NewLoggerWithCtx()
//	logger.Log(ctx, log.InfoLevel, "message with context")
//
// # Performance
//
// This library uses sync.Pool for entry reuse, inline optimizations,
// and zero-allocation fast paths for common operations, achieving
// sub-microsecond latency for most logging operations.
package log
