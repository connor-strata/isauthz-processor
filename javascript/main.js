#!/usr/bin/env node

const readline = require('readline');

/**
 * Process the input map and return authorization result
 * @param {Object} inputData - The input data as a key-value object
 * @returns {string} "authorized" or "unauthorized"
 */
function processAuthorizationRequest(inputData) {
    // Check if user is authenticated
    const authenticated = inputData['azure.authenticated'] || '';
    if (authenticated !== 'true') {
        return 'unauthorized';
    }

    // Get user attributes
    const role = inputData['azure.role'] || '';
    const department = inputData['azure.department'] || '';
    const groups = inputData['azure.groups'] || '';
    const email = inputData['azure.email'] || '';

    // Complex authorization logic based on multiple attributes

    // Admin users are always authorized
    if (role === 'admin') {
        return 'authorized';
    }

    // Engineering department developers are authorized
    if (department === 'Engineering' && groups.includes('developers')) {
        return 'authorized';
    }

    // Users with @example.com email and user role are authorized
    if (email.endsWith('@example.com') && role === 'user') {
        return 'authorized';
    }

    // Default to unauthorized
    return 'unauthorized';
}

/**
 * Main function that reads JSON from stdin line by line
 */
function main() {
    const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout,
        terminal: false
    });

    rl.on('line', (line) => {
        line = line.trim();

        // Skip empty lines
        if (!line) {
            return;
        }

        try {
            // Parse JSON input
            const inputData = JSON.parse(line);

            // Ensure it's an object
            if (typeof inputData !== 'object' || inputData === null || Array.isArray(inputData)) {
                console.error('Error parsing JSON: Input is not a JSON object');
                console.log('unauthorized');
                return;
            }

            // Convert all values to strings
            const stringData = {};
            for (const [key, value] of Object.entries(inputData)) {
                stringData[key] = String(value);
            }

            // Process the input and determine authorization
            const result = processAuthorizationRequest(stringData);

            // Output result
            console.log(result);

        } catch (error) {
            console.error(`Error parsing JSON: ${error.message}`);
            console.log('unauthorized');
        }
    });

    rl.on('close', () => {
        process.exit(0);
    });

    rl.on('error', (error) => {
        console.error(`Error reading from stdin: ${error.message}`);
        process.exit(1);
    });
}

if (require.main === module) {
    main();
}
