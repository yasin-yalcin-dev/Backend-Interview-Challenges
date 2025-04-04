# Observer Pattern

## Problem
Implement the Observer design pattern to create a newsletter subscription system. Users should be able to subscribe to different categories of news, and when new content is published, only subscribers of those categories should be notified.

## Requirements
1. Create a Subject interface that can add, remove, and notify observers
2. Create an Observer interface that can receive updates
3. Implement a Newsletter (Subject) that can publish content in different categories
4. Implement NewsletterSubscriber (Observer) that can subscribe to specific categories
5. When new content is published in a category, only subscribers of that category should be notified

## Examples
```go
// Create a newsletter
newsletter := NewNewsletter()

// Create subscribers
subscriber1 := NewNewsletterSubscriber("John")
subscriber2 := NewNewsletterSubscriber("Jane")
subscriber3 := NewNewsletterSubscriber("Bob")

// Subscribe to different categories
newsletter.AddObserver(subscriber1, "Technology")
newsletter.AddObserver(subscriber1, "Business")
newsletter.AddObserver(subscriber2, "Technology")
newsletter.AddObserver(subscriber3, "Sports")

// Publish content to a category
newsletter.PublishUpdate("Technology", "New AI breakthrough announced")
Expected Output
CopyNotifying subscribers of category: Technology
John received update for Technology: New AI breakthrough announced
Jane received update for Technology: New AI breakthrough announced