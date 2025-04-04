# Backend Interview Challenges

A comprehensive collection of backend coding challenges and solutions for technical interviews. This repository contains algorithmic problems, OOP design challenges, data structure implementations, and system design exercises commonly asked in backend developer interviews.

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Problem Categories](#problem-categories)
- [Running the Challenges](#running-the-challenges)
- [Contributing](#contributing)
- [License](#license)

## Overview

This repository aims to help backend developers prepare for technical interviews by providing a structured collection of coding challenges with well-documented solutions in Go. Each problem includes:

- Clear problem statement
- Example inputs and expected outputs
- Solution implementation
- Test cases

## Getting Started

### Prerequisites

- Go 1.24 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yasin-yalcin-dev/Backend-Interview-Challenges.git
cd Backend-Interview-Challenges

Build the project:

bashCopygo build -o interview-challenges ./cmd

Run the tool:

bashCopy./interview-challenges list
Project Structure
CopyBackend-Interview-Challenges/
├── cmd/                        # Command-line tool files
│   ├── main.go                 # Main entry point
│   ├── runner.go               # Problem runners
│   ├── listing.go              # Problem listing
│   └── usage.go                # Usage instructions
│
├── problems/                   # All problems organized by category
│   ├── algorithms/             # Algorithm problems
│   │   ├── stringreversal/     # Reverse a string
│   │   ├── palindrome/         # Check if a string is a palindrome
│   │   ├── countvowels/        # Count vowels in a string
│   │   ├── fizzbuzz/           # FizzBuzz implementation
│   │   ├── twosum/             # Find indices of two numbers that add up to target
│   │   └── firstrepeatingcharacter/ # Find first repeating character in a string
│   │
│   ├── datastructures/         # Data structure implementations (Coming soon)
│   ├── oop/                    # Object-oriented design problems (Coming soon)
│   └── systemdesign/           # System design challenges (Coming soon)
Problem Categories
Algorithms

String Reversal - Reverse a string
Palindrome Checker - Check if a string is a palindrome
Count Vowels - Count vowels in a string
FizzBuzz - Classic FizzBuzz implementation
Two Sum - Find indices of two numbers that add up to a target
First Repeating Character - Find the first repeating character in a string

OOP Design (Coming Soon)

Shape Hierarchy - Implement a polymorphic shape class hierarchy
Banking System - Design classes for a banking application
Factory Pattern - Implement payment methods using the Factory pattern
Observer Pattern - Create a newsletter subscription system
Singleton & Dependency Injection - Database connection pool implementation

Data Structures (Coming Soon)

Linked List - Implementation and operations
Stack & Queue - Implementation and applications
Binary Search Tree - Implementation and traversals
Graph - Representation and algorithms
Hash Table - Custom implementation

System Design (Coming Soon)

Rate Limiter - Design a rate limiting system
Microservice Communication - Design communication between microservices
Caching Layer - Implement a caching mechanism
URL Shortener - Design a URL shortening service

Running the Challenges
List all available problems:
bashCopy./interview-challenges list
Run a specific algorithm problem:
bashCopy./interview-challenges algorithms stringreversal "hello world"
./interview-challenges algorithms palindrome "racecar"
./interview-challenges algorithms countvowels "beautiful"
./interview-challenges algorithms fizzbuzz 15
./interview-challenges algorithms twosum "[2,7,11,15]" 9
./interview-challenges algorithms firstrepeatingcharacter "hello"
Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

Fork the repository
Create your feature branch (git checkout -b feature/amazing-feature)
Commit your changes (git commit -m 'Add some amazing feature')
Push to the branch (git push origin feature/amazing-feature)
Open a Pull Request