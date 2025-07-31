#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define MAX_LINE_LENGTH 4096
#define MAX_KEY_LENGTH 256
#define MAX_VALUE_LENGTH 256
#define MAX_PAIRS 32

typedef struct {
    char key[MAX_KEY_LENGTH];
    char value[MAX_VALUE_LENGTH];
} KeyValuePair;

typedef struct {
    KeyValuePair pairs[MAX_PAIRS];
    int count;
} InputData;

// Simple JSON parser for key-value pairs (very basic implementation)
bool parse_json_line(const char* line, InputData* data) {
    data->count = 0;
    
    // Skip whitespace and opening brace
    const char* ptr = line;
    while (*ptr == ' ' || *ptr == '\t' || *ptr == '\n') ptr++;
    if (*ptr != '{') return false;
    ptr++;
    
    while (*ptr && *ptr != '}' && data->count < MAX_PAIRS) {
        // Skip whitespace
        while (*ptr == ' ' || *ptr == '\t' || *ptr == '\n' || *ptr == ',') ptr++;
        if (*ptr == '}') break;
        
        // Parse key
        if (*ptr != '"') return false;
        ptr++; // skip opening quote
        
        int key_len = 0;
        while (*ptr && *ptr != '"' && key_len < MAX_KEY_LENGTH - 1) {
            data->pairs[data->count].key[key_len++] = *ptr++;
        }
        data->pairs[data->count].key[key_len] = '\0';
        
        if (*ptr != '"') return false;
        ptr++; // skip closing quote
        
        // Skip whitespace and colon
        while (*ptr == ' ' || *ptr == '\t') ptr++;
        if (*ptr != ':') return false;
        ptr++;
        while (*ptr == ' ' || *ptr == '\t') ptr++;
        
        // Parse value
        if (*ptr != '"') return false;
        ptr++; // skip opening quote
        
        int value_len = 0;
        while (*ptr && *ptr != '"' && value_len < MAX_VALUE_LENGTH - 1) {
            data->pairs[data->count].value[value_len++] = *ptr++;
        }
        data->pairs[data->count].value[value_len] = '\0';
        
        if (*ptr != '"') return false;
        ptr++; // skip closing quote
        
        data->count++;
    }
    
    return true;
}

// Get value by key from InputData
const char* get_value(const InputData* data, const char* key) {
    for (int i = 0; i < data->count; i++) {
        if (strcmp(data->pairs[i].key, key) == 0) {
            return data->pairs[i].value;
        }
    }
    return "";
}

// Check if a string contains a substring
bool contains(const char* haystack, const char* needle) {
    return strstr(haystack, needle) != NULL;
}

// Check if a string ends with a suffix
bool ends_with(const char* str, const char* suffix) {
    size_t str_len = strlen(str);
    size_t suffix_len = strlen(suffix);
    
    if (suffix_len > str_len) return false;
    
    return strcmp(str + str_len - suffix_len, suffix) == 0;
}

// Process the input data and return authorization result
const char* process_authorization_request(const InputData* input) {
    // Check if user is authenticated
    const char* authenticated = get_value(input, "azure.authenticated");
    if (strcmp(authenticated, "true") != 0) {
        return "unauthorized";
    }
    
    // Get user attributes
    const char* role = get_value(input, "azure.role");
    const char* department = get_value(input, "azure.department");
    const char* groups = get_value(input, "azure.groups");
    const char* email = get_value(input, "azure.email");
    
    // Complex authorization logic based on multiple attributes
    
    // Admin users are always authorized
    if (strcmp(role, "admin") == 0) {
        return "authorized";
    }
    
    // Engineering department developers are authorized
    if (strcmp(department, "Engineering") == 0 && contains(groups, "developers")) {
        return "authorized";
    }
    
    // Users with @example.com email and user role are authorized
    if (ends_with(email, "@example.com") && strcmp(role, "user") == 0) {
        return "authorized";
    }
    
    // Default to unauthorized
    return "unauthorized";
}

int main() {
    char line[MAX_LINE_LENGTH];
    InputData input_data;
    
    // Read input line by line from stdin
    while (fgets(line, sizeof(line), stdin)) {
        // Remove newline character
        line[strcspn(line, "\n")] = '\0';
        
        // Skip empty lines
        if (strlen(line) == 0) {
            continue;
        }
        
        // Parse JSON input
        if (parse_json_line(line, &input_data)) {
            // Process the input and determine authorization
            const char* result = process_authorization_request(&input_data);
            
            // Output result
            printf("%s\n", result);
        } else {
            fprintf(stderr, "Error parsing JSON: %s\n", line);
            printf("unauthorized\n");
        }
    }
    
    return 0;
}
