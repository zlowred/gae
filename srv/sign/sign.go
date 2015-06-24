package sign

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"net/http"
	"time"

	"github.com/zlowred/gae/srv/mdl"
)

// [START func_sign]
func Sign(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	g := mdl.Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Greeting", mdl.GuestbookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// [END func_sign]
