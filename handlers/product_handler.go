package handlers

import (
	"fmt"
	"github.com/yigitsadic/qrmenum_client/client"
	"github.com/yigitsadic/qrmenum_client/store"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// Renders not found response.
func renderNotFoundResponse(w http.ResponseWriter, tmpl *template.Template) {
	w.WriteHeader(http.StatusNotFound)
	err := tmpl.Execute(w, nil)

	if err != nil {
		_, _ = fmt.Fprint(w, "Unable to write template")
	}
}

// Renders products.
func renderProducts(w http.ResponseWriter, tmpl *template.Template, data []client.ProductResponse) {
	w.WriteHeader(http.StatusOK)

	err := tmpl.Execute(w, map[string][]client.ProductResponse{
		"Products": data,
	})
	if err != nil {
		_, _ = fmt.Fprintf(w, "Unable to write template")
	}
}

func ProductHandler(s store.Store, mu *sync.Mutex, show *template.Template, notFound *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		q := r.URL.Query().Get("q")
		if q == "" {
			// Return not found response if query param not present.

			log.Println("q not found in request")
			renderNotFoundResponse(w, notFound)
			return
		}

		p, err := s.GetMapItem(q)
		if err != nil {
			log.Printf("Requested %q key not found in in-memory store\n", q)

			c := s.GetClient()
			resp, err := c.FetchFromCMS(q)
			if err != nil {
				log.Printf("Queried for %q in CMS, but found nothing\n", q)

				s.SetMapItem(q, nil)

				renderNotFoundResponse(w, notFound)
				return
			}

			s.SetMapItem(q, resp)
			log.Printf("Fetched and assigned new content from CMS for %q\n", q)

			renderProducts(w, show, resp)
			return
		}

		if p.IsExpired() {
			// Re-fetch from CMS.

			log.Printf("Given key %q is expired will re-fetch\n", q)

			c := s.GetClient()
			resp, err := c.FetchFromCMS(q)
			if err != nil {
				log.Printf("Tried to fetch %q from CMS but found nothing\n", q)

				s.SetMapItem(q, nil)

				renderNotFoundResponse(w, notFound)
				return
			}

			s.SetMapItem(q, resp)
			log.Printf("Fetched and assigned new content from CMS for %q\n", q)

			renderProducts(w, show, resp)
			return
		} else {
			// Render products.

			log.Printf("Found %q in in-memory store\n", q)
			renderProducts(w, show, p.Products)
			return
		}
	}
}
