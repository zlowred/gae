package mdl

import (
	"appengine"
	"appengine/datastore"
	"time"
)

// [START greeting_struct]
type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

// [END greeting_struct]

// guestbookKey returns the key used for all guestbook entries.
func GuestbookKey(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}
