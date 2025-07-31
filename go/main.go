package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

// Pre-computed constants to avoid string allocations
const (
	authorized   = "authorized"
	unauthorized = "unauthorized"
	trueValue    = "true"
	adminRole    = "admin" 
	userRole     = "user"
	engineering  = "Engineering"
	developers   = "developers"
	exampleDomain = "@example.com"
)

func main() {
	// Create larger buffered reader for stdin (default is 4KB, increase to 64KB)
	reader := bufio.NewReaderSize(os.Stdin, 65536)
	
	// Create buffered writer for stdout with optimal buffer size
	writer := bufio.NewWriterSize(os.Stdout, 8192)
	defer writer.Flush()

	// Pre-allocate byte slice to avoid repeated allocations
	var line []byte
	var err error

	// Use ReadBytes instead of Scanner for better performance
	for {
		line, err = reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Stderr.WriteString("Error reading from stdin: ")
			os.Stderr.WriteString(err.Error())
			os.Stderr.WriteString("\n")
			os.Exit(1)
		}

		// Remove newline character efficiently
		if len(line) > 0 && line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Parse JSON input directly from byte slice (avoid string conversion)
		var input map[string]string
		if err := json.Unmarshal(line, &input); err != nil {
			// Optimized error writing to stderr
			os.Stderr.WriteString("Error parsing JSON: ")
			os.Stderr.WriteString(err.Error())
			os.Stderr.WriteString("\n")

			// Write "unauthorized" using bufio with constant
			writer.WriteString(unauthorized)
			writer.WriteByte('\n')
			writer.Flush() // Flush for error case
			continue
		}

		// Process the input and determine authorization
		result := processAuthorizationRequest(input)

		// Write result efficiently using constants
		writer.WriteString(result)
		writer.WriteByte('\n')
		writer.Flush() // Flush immediately for real-time output
	}
}

// processAuthorizationRequest processes the input map and returns authorization result
func processAuthorizationRequest(input map[string]string) string {
	// Check if user is authenticated using pre-computed constant
	authenticated, exists := input["azure.authenticated"]
	if !exists || authenticated != trueValue {
		return unauthorized
	}

	// Get user attributes
	role := input["azure.role"]
	department := input["azure.department"]
	groups := input["azure.groups"]
	email := input["azure.email"]

	// Complex authorization logic based on multiple attributes

	// Admin users are always authorized
	if role == adminRole {
		return authorized
	}

	// Engineering department developers are authorized
	if department == engineering && strings.Contains(groups, developers) {
		return authorized
	}

	// Users with @example.com email and user role are authorized
	if strings.HasSuffix(email, exampleDomain) && role == userRole {
		return authorized
	}

	// Default to unauthorized
	return unauthorized
}
