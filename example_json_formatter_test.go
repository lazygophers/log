package log

import (
	"fmt"
	"os"
)

// Example_jsonFormatter demonstrates JSON logging usage
func Example_jsonFormatter() {
	// Create logger with JSON output
	logger := New()
	logger.Format = &JSONFormatter{}
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	// Basic JSON logging
	logger.Info("Application started in JSON format")
	logger.Error("Error occurred in JSON format")

	// Pretty print JSON
	logger.Format = &JSONFormatter{EnablePrettyPrint: true}
	logger.Warn("This is pretty printed JSON")

	// Minimal JSON (no caller, no trace)
	logger.Format = &JSONFormatter{
		DisableCaller: true,
		DisableTrace:  true,
	}
	logger.Debug("Minimal JSON log")

	// Output to file with JSON format
	// logger.SetOutput(GetOutputWriterHourly("./logs/app.log"))
	// logger.Format = &JSONFormatter{}
	// logger.Info("JSON logs to file")
}

// Example_jsonFormatter_multipleOutputs shows logging to both console (text) and file (JSON)
func Example_jsonFormatter_multipleOutputs() {
	// Console logger with text format
	consoleLogger := New()
	consoleLogger.SetOutput(os.Stdout)

	// File logger with JSON format
	// fileLogger := New()
	// fileLogger.SetOutput(GetOutputWriterHourly("./logs/app.log"))
	// fileLogger.Format = &JSONFormatter{}

	// Use both loggers as needed
	consoleLogger.Info("This goes to console as text")
	// fileLogger.Info("This goes to file as JSON")
}

func Example_jsonFormatter_withContext() {
	// This requires the logctx subpackage
	// import logctx "github.com/lazygophers/log/logctx"
	//
	// logger := logctx.New()
	// logger.Format = &log.JSONFormatter{}
	// logger.Info(ctx, "JSON with context")
	_ = fmt.Sprintf("context example")
}

// Example_jsonFormatter_fields shows how to log structured data
func Example_jsonFormatter_fields() {
	logger := New()
	logger.Format = &JSONFormatter{}
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	// For structured logging, include fields in the message
	// Note: This is a simple approach. For true structured logging,
	// consider using the Field API when implemented.

	// Simple structured logging
	userID := "user-123"
	action := "login"
	logger.Infof("User %s performed action: %s", userID, action)

	// Or build JSON manually
	// message := fmt.Sprintf(`{"user_id":"%s","action":"%s"}`, userID, action)
	// logger.Info(message)
}
