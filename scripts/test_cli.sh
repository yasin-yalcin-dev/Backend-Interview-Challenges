#!/bin/bash

# Script to test the CLI functionality

echo "Testing CLI functionality..."
echo "==========================="

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Build the CLI
echo -e "\n${GREEN}Building CLI...${NC}"
go build -o interview-challenges ./cmd

# Track failures
failures=0
total=0

# Function to test a CLI command
test_command() {
    local cmd=$1
    local description=$2
    
    echo -e "\n${GREEN}Testing:${NC} $description"
    echo "Command: ./interview-challenges $cmd"
    
    # Run the command
    ./interview-challenges $cmd
    
    # Check if command succeeded
    if [ $? -ne 0 ]; then
        echo -e "${RED}Failed${NC}: $description"
        ((failures++))
    else
        echo -e "${GREEN}Passed${NC}: $description"
    fi
    
    ((total++))
}

# Test listing problems
test_command "list" "List all problems"

# Test algorithm problems
test_command "algorithms stringreversal \"hello world\"" "String Reversal Algorithm"
test_command "algorithms palindrome \"racecar\"" "Palindrome Algorithm"
test_command "algorithms countvowels \"beautiful\"" "Count Vowels Algorithm"
test_command "algorithms fizzbuzz 15" "FizzBuzz Algorithm"
test_command "algorithms twosum \"[2,7,11,15]\" 9" "Two Sum Algorithm"
test_command "algorithms firstrepeatingcharacter \"hello\"" "First Repeating Character Algorithm"

# Test OOP problems
test_command "oop shapehierarchy" "Shape Hierarchy OOP Design"
test_command "oop bankingsystem" "Banking System OOP Design"
test_command "oop factorypattern" "Factory Pattern OOP Design"
test_command "oop observerpattern" "Observer Pattern OOP Design"
test_command "oop singleton" "Singleton Pattern OOP Design"

# Test data structure problems
test_command "datastructures linkedlist" "Linked List Data Structure"
test_command "datastructures stackqueue" "Stack & Queue Data Structures"
test_command "datastructures binarysearchtree" "Binary Search Tree Data Structure"
test_command "datastructures hashtable" "Hash Table Data Structure"
test_command "datastructures graph" "Graph Data Structure"

# Report summary
echo -e "\n==========================="
echo -e "Test Summary: ${total} commands run, ${failures} failures"

if [ "$failures" -eq 0 ]; then
    echo -e "${GREEN}All CLI tests passed successfully!${NC}"
    
    # Clean up
    echo -e "\nCleaning up..."
    rm interview-challenges
    
    exit 0
else
    echo -e "${RED}Some CLI tests failed. Please check the output above.${NC}"
    exit 1
fi