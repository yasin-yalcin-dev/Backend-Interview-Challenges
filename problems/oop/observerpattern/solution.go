/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package observerpattern

import "fmt"

// RunExample demonstrates the observer pattern implementation
func RunExample() {
	fmt.Println("Observer Pattern Example:")
	fmt.Println("------------------------")

	// Create a newsletter
	newsletter := NewNewsletter()

	// Create subscribers
	subscriber1 := NewNewsletterSubscriber("John")
	subscriber2 := NewNewsletterSubscriber("Jane")
	subscriber3 := NewNewsletterSubscriber("Bob")

	fmt.Println("Subscribing users to different categories...")

	// Subscribe to different categories
	newsletter.AddObserver(subscriber1, "Technology")
	newsletter.AddObserver(subscriber1, "Business")
	newsletter.AddObserver(subscriber2, "Technology")
	newsletter.AddObserver(subscriber3, "Sports")

	// Show subscription status
	fmt.Printf("\nSubscriber counts:\n")
	fmt.Printf("Technology: %d subscribers\n", newsletter.GetSubscribersCount("Technology"))
	fmt.Printf("Business: %d subscribers\n", newsletter.GetSubscribersCount("Business"))
	fmt.Printf("Sports: %d subscribers\n", newsletter.GetSubscribersCount("Sports"))

	// Publish content to different categories
	fmt.Printf("\nPublishing updates to different categories:\n\n")

	newsletter.PublishUpdate("Technology", "New AI breakthrough announced")
	fmt.Println()

	newsletter.PublishUpdate("Business", "Stock market reaches record high")
	fmt.Println()

	newsletter.PublishUpdate("Sports", "Local team wins championship")
	fmt.Println()

	// Unsubscribe and publish again
	fmt.Println("Unsubscribing Jane from Technology...")
	newsletter.RemoveObserver(subscriber2, "Technology")

	fmt.Printf("\nTechnology now has %d subscribers\n\n", newsletter.GetSubscribersCount("Technology"))

	newsletter.PublishUpdate("Technology", "Tech conference announced for next month")
}
