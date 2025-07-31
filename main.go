package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read input line by line from stdin
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines
		if line == "" {
			continue
		}

		// Parse JSON input
		var input map[string]string
		if err := json.Unmarshal([]byte(line), &input); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			fmt.Println("unauthorized")
			continue
		}

		// Process the input and determine authorization
		result := processAuthorizationRequest(input)

		// Output result
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
}

// processAuthorizationRequest processes the input map and returns authorization result
func processAuthorizationRequest(input map[string]string) string {
	// TODO: Implement authorization logic here
	// This is where you would add your specific authorization processing

	// Placeholder logic - you can replace this with your actual authorization logic
	// For now, we'll just check if certain required fields are present

	// Example: Check if required fields exist
	if _, hasUser := input["user"]; !hasUser {
		return "unauthorized"
	}

	if _, hasResource := input["resource"]; !hasResource {
		return "unauthorized"
	}

	if _, hasAction := input["action"]; !hasAction {
		return "unauthorized"
	}

	// TODO: Add your actual authorization logic here
	// This could involve:
	// - Checking permissions in a database
	// - Validating JWT tokens
	// - Consulting policy engines
	// - Checking role-based access control
	// - etc.

	// For now, return authorized if all required fields are present
	return "authorized"
}
