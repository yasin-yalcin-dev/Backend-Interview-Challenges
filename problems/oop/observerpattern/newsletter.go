/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package observerpattern

import (
	"fmt"
)

// Newsletter implements the Subject interface
type Newsletter struct {
	// Map of category -> slice of observers
	subscribers map[string][]Observer
}

// NewNewsletter creates a new newsletter with no subscribers
func NewNewsletter() *Newsletter {
	return &Newsletter{
		subscribers: make(map[string][]Observer),
	}
}

// AddObserver adds a subscriber to a specific category
func (n *Newsletter) AddObserver(observer Observer, category string) {
	// If this is the first subscriber for this category, initialize the slice
	if _, exists := n.subscribers[category]; !exists {
		n.subscribers[category] = []Observer{}
	}

	// Check if the observer is already subscribed to this category
	for _, existingObserver := range n.subscribers[category] {
		if existingObserver == observer {
			return // Already subscribed
		}
	}

	// Add the observer to the category
	n.subscribers[category] = append(n.subscribers[category], observer)
}

// RemoveObserver removes a subscriber from a specific category
func (n *Newsletter) RemoveObserver(observer Observer, category string) {
	observers, exists := n.subscribers[category]
	if !exists {
		return // No subscribers for this category
	}

	// Filter out the observer to remove
	updatedObservers := []Observer{}
	for _, existingObserver := range observers {
		if existingObserver != observer {
			updatedObservers = append(updatedObservers, existingObserver)
		}
	}

	n.subscribers[category] = updatedObservers
}

// NotifyObservers sends an update to all subscribers of a specific category
func (n *Newsletter) NotifyObservers(category, message string) {
	observers, exists := n.subscribers[category]
	if !exists || len(observers) == 0 {
		fmt.Printf("No subscribers for category: %s\n", category)
		return
	}

	fmt.Printf("Notifying subscribers of category: %s\n", category)
	for _, observer := range observers {
		observer.Update(category, message)
	}
}

// PublishUpdate publishes a new update to a specific category
func (n *Newsletter) PublishUpdate(category, message string) {
	fmt.Printf("Publishing to category: %s\n", category)
	n.NotifyObservers(category, message)
}

// GetCategories returns a list of all categories that have subscribers
func (n *Newsletter) GetCategories() []string {
	categories := make([]string, 0, len(n.subscribers))
	for category := range n.subscribers {
		categories = append(categories, category)
	}
	return categories
}

// GetSubscribersCount returns the number of subscribers for a specific category
func (n *Newsletter) GetSubscribersCount(category string) int {
	observers, exists := n.subscribers[category]
	if !exists {
		return 0
	}
	return len(observers)
}

// NewsletterSubscriber implements the Observer interface
type NewsletterSubscriber struct {
	name string
}

// NewNewsletterSubscriber creates a new newsletter subscriber
func NewNewsletterSubscriber(name string) *NewsletterSubscriber {
	return &NewsletterSubscriber{
		name: name,
	}
}

// Update receives an update from the subject
func (s *NewsletterSubscriber) Update(category, message string) {
	fmt.Printf("%s received update for %s: %s\n", s.name, category, message)
}

// GetName returns the subscriber's name
func (s *NewsletterSubscriber) GetName() string {
	return s.name
}
