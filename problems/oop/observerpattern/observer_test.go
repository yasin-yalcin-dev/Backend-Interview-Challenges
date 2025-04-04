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
	"testing"
)

// MockObserver implements the Observer interface for testing
type MockObserver struct {
	name            string
	updatesReceived map[string][]string // Category -> slice of messages
	updateCallCount int
}

// NewMockObserver creates a new mock observer for testing
func NewMockObserver(name string) *MockObserver {
	return &MockObserver{
		name:            name,
		updatesReceived: make(map[string][]string),
		updateCallCount: 0,
	}
}

// Update implementation for the mock observer
func (m *MockObserver) Update(category, message string) {
	m.updateCallCount++

	if _, exists := m.updatesReceived[category]; !exists {
		m.updatesReceived[category] = []string{}
	}

	m.updatesReceived[category] = append(m.updatesReceived[category], message)
}

// GetName returns the observer's name
func (m *MockObserver) GetName() string {
	return m.name
}

// TestNewsletter tests the newsletter implementation
func TestNewsletter(t *testing.T) {
	newsletter := NewNewsletter()

	if len(newsletter.GetCategories()) != 0 {
		t.Errorf("Expected 0 categories, got %d", len(newsletter.GetCategories()))
	}

	observer1 := NewMockObserver("Observer1")
	observer2 := NewMockObserver("Observer2")

	// Test adding observers
	newsletter.AddObserver(observer1, "Technology")
	newsletter.AddObserver(observer2, "Technology")
	newsletter.AddObserver(observer1, "Business")

	if newsletter.GetSubscribersCount("Technology") != 2 {
		t.Errorf("Expected 2 subscribers for Technology, got %d", newsletter.GetSubscribersCount("Technology"))
	}

	if newsletter.GetSubscribersCount("Business") != 1 {
		t.Errorf("Expected 1 subscriber for Business, got %d", newsletter.GetSubscribersCount("Business"))
	}

	// Test notifications
	newsletter.NotifyObservers("Technology", "Test message")

	if observer1.updateCallCount != 1 {
		t.Errorf("Expected 1 update for observer1, got %d", observer1.updateCallCount)
	}

	if observer2.updateCallCount != 1 {
		t.Errorf("Expected 1 update for observer2, got %d", observer2.updateCallCount)
	}

	// Test removing observer
	newsletter.RemoveObserver(observer1, "Technology")

	if newsletter.GetSubscribersCount("Technology") != 1 {
		t.Errorf("Expected 1 subscriber for Technology after removal, got %d",
			newsletter.GetSubscribersCount("Technology"))
	}

	// Test notification after removal
	newsletter.NotifyObservers("Technology", "Second message")

	if observer1.updateCallCount != 1 {
		t.Errorf("Expected still 1 update for observer1 after removal, got %d", observer1.updateCallCount)
	}

	if observer2.updateCallCount != 2 {
		t.Errorf("Expected 2 updates for observer2, got %d", observer2.updateCallCount)
	}

	// Check if correct messages were received
	techMessages, exists := observer2.updatesReceived["Technology"]
	if !exists || len(techMessages) != 2 {
		t.Errorf("Expected 2 Technology messages for observer2")
	}

	if len(techMessages) >= 2 && techMessages[1] != "Second message" {
		t.Errorf("Expected 'Second message', got '%s'", techMessages[1])
	}
}
