use std::collections::HashMap;
use std::io::{self, BufRead, BufReader};

/// Process the input map and return authorization result
fn process_authorization_request(input: &HashMap<String, String>) -> &'static str {
    // Check if user is authenticated
    let authenticated = input
        .get("azure.authenticated")
        .map(|s| s.as_str())
        .unwrap_or("");
    if authenticated != "true" {
        return "unauthorized";
    }

    // Get user attributes
    let role = input.get("azure.role").map(|s| s.as_str()).unwrap_or("");
    let department = input
        .get("azure.department")
        .map(|s| s.as_str())
        .unwrap_or("");
    let groups = input.get("azure.groups").map(|s| s.as_str()).unwrap_or("");
    let email = input.get("azure.email").map(|s| s.as_str()).unwrap_or("");

    // Complex authorization logic based on multiple attributes

    // Admin users are always authorized
    if role == "admin" {
        return "authorized";
    }

    // Engineering department developers are authorized
    if department == "Engineering" && groups.contains("developers") {
        return "authorized";
    }

    // Users with @example.com email and user role are authorized
    if email.ends_with("@example.com") && role == "user" {
        return "authorized";
    }

    // Default to unauthorized
    "unauthorized"
}

fn main() -> io::Result<()> {
    let stdin = io::stdin();
    let reader = BufReader::new(stdin.lock());

    // Read input line by line from stdin
    for line in reader.lines() {
        let line = line?;
        let line = line.trim();

        // Skip empty lines
        if line.is_empty() {
            continue;
        }

        // Parse JSON input
        match serde_json::from_str::<HashMap<String, serde_json::Value>>(line) {
            Ok(json_data) => {
                // Convert all values to strings
                let mut input_data = HashMap::new();
                for (key, value) in json_data {
                    let string_value = match value {
                        serde_json::Value::String(s) => s,
                        serde_json::Value::Bool(b) => b.to_string(),
                        serde_json::Value::Number(n) => n.to_string(),
                        _ => value.to_string().trim_matches('"').to_string(),
                    };
                    input_data.insert(key, string_value);
                }

                // Process the input and determine authorization
                let result = process_authorization_request(&input_data);

                // Output result
                println!("{}", result);
            }
            Err(e) => {
                eprintln!("Error parsing JSON: {}", e);
                println!("unauthorized");
            }
        }
    }

    Ok(())
}
