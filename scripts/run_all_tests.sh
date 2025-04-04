#!/bin/bash

# Script to run all tests in the project

echo "Running all tests in the project..."
echo "=================================="

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Array of categories
categories=("algorithms" "oop" "datastructures")

# Track failures
failures=0
total=0

# Function to run tests for a specific directory
run_tests() {
    local dir=$1
    local category=$2
    
    echo -e "\n${GREEN}Testing $category:${NC} $dir"
    
    # Go to directory and run tests
    cd "$dir" || { echo "Failed to change to directory $dir"; return 1; }
    
    # Run tests with verbose flag and capture output
    go test -v
    
    # Check if tests failed
    if [ $? -ne 0 ]; then
        ((failures++))
    fi
    
    ((total++))
    
    # Go back to the root directory
    cd - > /dev/null
}

# Get the project root directory (assuming the script is in the project root)
root_dir=$(pwd)

# Run tests for each category
for category in "${categories[@]}"; do
    # Find all test directories for this category
    find "$root_dir/problems/$category" -type d -name "*" -not -path "*/\.*" | while read -r dir; do
        # Check if directory contains test files
        if ls "$dir"/*_test.go 1> /dev/null 2>&1; then
            run_tests "$dir" "$category"
        fi
    done
done

# Report summary
echo -e "\n=================================="
echo -e "Test Summary: ${total} test suites run, ${failures} failures"

if [ "$failures" -eq 0 ]; then
    echo -e "${GREEN}All tests passed successfully!${NC}"
    exit 0
else
    echo -e "${RED}Some tests failed. Please check the output above.${NC}"
    exit 1
fi