package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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
	// Check if user is authenticated
	authenticated, exists := input["azure.authenticated"]
	if !exists || authenticated != "true" {
		return "unauthorized"
	}

	// Get user attributes
	role := input["azure.role"]
	department := input["azure.department"]
	groups := input["azure.groups"]
	email := input["azure.email"]

	// Complex authorization logic based on multiple attributes

	// Admin users are always authorized
	if role == "admin" {
		return "authorized"
	}

	// Engineering department developers are authorized
	if department == "Engineering" && strings.Contains(groups, "developers") {
		return "authorized"
	}

	// Users with @example.com email and user role are authorized
	if strings.HasSuffix(email, "@example.com") && role == "user" {
		return "authorized"
	}

	// Default to unauthorized
	return "unauthorized"
}
