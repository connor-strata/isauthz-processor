#!/usr/bin/env python3

import json
import sys
import logging

def process_authorization_request(input_data):
    """Process the input map and return authorization result"""
    
    # Check if user is authenticated
    authenticated = input_data.get("azure.authenticated", "")
    if authenticated != "true":
        return "unauthorized"
    
    # Get user attributes
    role = input_data.get("azure.role", "")
    department = input_data.get("azure.department", "")
    groups = input_data.get("azure.groups", "")
    email = input_data.get("azure.email", "")
    
    # Complex authorization logic based on multiple attributes
    
    # Admin users are always authorized
    if role == "admin":
        return "authorized"
    
    # Engineering department developers are authorized
    if department == "Engineering" and "developers" in groups:
        return "authorized"
    
    # Users with @example.com email and user role are authorized
    if email.endswith("@example.com") and role == "user":
        return "authorized"
    
    # Default to unauthorized
    return "unauthorized"

def main():
    """Main function that reads JSON from stdin line by line"""
    
    # Configure logging to stderr
    logging.basicConfig(stream=sys.stderr, level=logging.ERROR, 
                       format='%(asctime)s %(levelname)s: %(message)s')
    
    try:
        # Read input line by line from stdin
        for line in sys.stdin:
            line = line.strip()
            
            # Skip empty lines
            if not line:
                continue
            
            try:
                # Parse JSON input
                input_data = json.loads(line)
                
                # Ensure it's a dictionary with string values
                if not isinstance(input_data, dict):
                    logging.error("Input is not a JSON object")
                    print("unauthorized")
                    continue
                
                # Convert all values to strings
                input_data = {k: str(v) for k, v in input_data.items()}
                
                # Process the input and determine authorization
                result = process_authorization_request(input_data)
                
                # Output result
                print(result)
                
            except json.JSONDecodeError as e:
                logging.error(f"Error parsing JSON: {e}")
                print("unauthorized")
                continue
                
    except KeyboardInterrupt:
        pass
    except Exception as e:
        logging.error(f"Unexpected error: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()
