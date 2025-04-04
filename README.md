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
```

2. Build the project:
```bash
go build -o interview-challenges ./cmd
```

3. Run the tool:
```bash
./interview-challenges list
```

## Project Structure

```
Backend-Interview-Challenges/
├── cmd/                        # Command-line tool files
│   ├── main.go                 # Main entry point
│   ├── runner.go               # Algorithm problem runners
│   ├── runner_oop.go           # OOP problem runners
│   ├── runner_datastructures.go # Data structure problem runners
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
│   ├── oop/                    # Object-oriented design problems
│   │   ├── shapehierarchy/     # Polymorphic shape class hierarchy
│   │   ├── bankingsystem/      # Banking application class design
│   │   ├── factorypattern/     # Factory pattern for payment methods
│   │   ├── observerpattern/    # Observer pattern for newsletter system
│   │   └── singleton/          # Singleton pattern with dependency injection
│   │
│   ├── datastructures/         # Data structure implementations
│   │   ├── linkedlist/         # Linked list implementation
│   │   ├── stackqueue/         # Stack and queue implementations
│   │   ├── binarysearchtree/   # Binary search tree implementation
│   │   ├── hashtable/          # Hash table implementation
│   │   └── graph/              # Graph implementation with algorithms
│   │
│   └── systemdesign/           # System design challenges (Coming soon)
```

## Problem Categories

### Algorithms
1. **String Reversal** - Reverse a string
2. **Palindrome Checker** - Check if a string is a palindrome
3. **Count Vowels** - Count vowels in a string
4. **FizzBuzz** - Classic FizzBuzz implementation
5. **Two Sum** - Find indices of two numbers that add up to a target
6. **First Repeating Character** - Find the first repeating character in a string

### OOP Design
1. **Shape Hierarchy** - Implement a polymorphic shape class hierarchy
2. **Banking System** - Design classes for a banking application
3. **Factory Pattern** - Implement payment methods using the Factory pattern
4. **Observer Pattern** - Create a newsletter subscription system
5. **Singleton & Dependency Injection** - Database connection pool implementation

### Data Structures
1. **Linked List** - Implementation and basic operations
2. **Stack & Queue** - Implementation with array and linked list approaches
3. **Binary Search Tree** - Implementation with traversal and search algorithms
4. **Hash Table** - Implementation with collision handling
5. **Graph** - Implementation with BFS, DFS, and shortest path algorithms

### System Design (Coming Soon)
- Rate Limiter - Design a rate limiting system
- Microservice Communication - Design communication between microservices
- Caching Layer - Implement a caching mechanism
- URL Shortener - Design a URL shortening service

## Running the Challenges

List all available problems:
```bash
./interview-challenges list
```

Run a specific algorithm problem:
```bash
./interview-challenges algorithms stringreversal "hello world"
./interview-challenges algorithms palindrome "racecar"
./interview-challenges algorithms countvowels "beautiful"
./interview-challenges algorithms fizzbuzz 15
./interview-challenges algorithms twosum "[2,7,11,15]" 9
./interview-challenges algorithms firstrepeatingcharacter "hello"
```

Run a specific OOP design problem:
```bash
./interview-challenges oop shapehierarchy
./interview-challenges oop bankingsystem
./interview-challenges oop factorypattern
./interview-challenges oop observerpattern
./interview-challenges oop singleton
```

Run a specific data structure implementation:
```bash
./interview-challenges datastructures linkedlist
./interview-challenges datastructures stackqueue
./interview-challenges datastructures binarysearchtree
./interview-challenges datastructures hashtable
./interview-challenges datastructures graph
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
