package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func isAuthorized(input map[string]string) bool {
	// Check authentication
	if input["azure.authenticated"] != "true" {
		return false
	}

	role := input["azure.role"]

	// Admin check
	if role == "admin" {
		return true
	}

	// Engineering developer check
	if input["azure.department"] == "Engineering" && strings.Contains(input["azure.groups"], "developers") {
		return true
	}

	// Example.com user check
	if role == "user" && strings.HasSuffix(input["azure.email"], "@example.com") {
		return true
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		var input map[string]string
		if err := json.Unmarshal([]byte(line), &input); err != nil {
			fmt.Println("unauthorized")
			continue
		}

		if isAuthorized(input) {
			fmt.Println("authorized")
		} else {
			fmt.Println("unauthorized")
		}
	}
}
