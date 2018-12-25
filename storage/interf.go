package storage

import (
	"net/url"
)

// Storage is an interface which handles Collector's internal data,
// like visited urls and cookies.
// The default Storage of the Collector is the InMemoryStorage.
// Collector's storage can be changed by calling Collector.SetStorage()
// function.
type Storage interface {
	// Init initializes the storage
	Init() error
	// Visited receives and stores a request ID that is visited by the Collector
	Visited(requestID uint64) error
	// IsVisited returns true if the request was visited before IsVisited
	// is called
	IsVisited(requestID uint64) (bool, error)
	// Cookies retrieves stored cookies for a given host
	Cookies(u *url.URL) string
	// SetCookies stores cookies for a given host
	SetCookies(u *url.URL, cookies string)
}
