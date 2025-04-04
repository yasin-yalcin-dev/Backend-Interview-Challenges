/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

// problems/oop/observerpattern/observer.go
package observerpattern

// Observer interface defines the method that will be called when the subject changes
type Observer interface {
	Update(category, message string)
	GetName() string
}

// Subject interface defines methods to manage observers and notify them of changes
type Subject interface {
	AddObserver(observer Observer, category string)
	RemoveObserver(observer Observer, category string)
	NotifyObservers(category, message string)
}
