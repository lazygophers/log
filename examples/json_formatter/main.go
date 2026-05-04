package main

import (
	"fmt"

	"github.com/lazygophers/log"
)

func main() {
	fmt.Println("=== JSON Formatter Demo ===\n")

	// 1. Basic JSON logging
	fmt.Println("1. Basic JSON Logging:")
	logger := log.New()
	logger.Format = &log.JSONFormatter{}
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Info("Application started")
	logger.Error("An error occurred")

	fmt.Println()

	// 2. Pretty print JSON
	fmt.Println("2. Pretty Print JSON:")
	logger.Format = &log.JSONFormatter{EnablePrettyPrint: true}
	logger.Warn("Warning with pretty print")

	fmt.Println()

	// 3. Minimal JSON
	fmt.Println("3. Minimal JSON (no caller/trace):")
	logger.Format = &log.JSONFormatter{
		DisableCaller: true,
		DisableTrace:  true,
	}
	logger.Debug("Minimal debug log")

	fmt.Println()

	// 4. JSON with all fields
	fmt.Println("4. JSON with All Fields:")
	logger.Format = &log.JSONFormatter{}
	logger.EnableCaller(true)
	logger.EnableTrace(true)
	log.SetTrace("trace-12345")
	logger.Info("Full JSON log with all metadata")

	fmt.Println("\n=== Demo Complete ===")
}
