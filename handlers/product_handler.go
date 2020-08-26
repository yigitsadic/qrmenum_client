package handlers

import (
	"fmt"
	"github.com/yigitsadic/qrmenum_client/client"
	"github.com/yigitsadic/qrmenum_client/store"
	"html/template"
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

func renderProducts(w http.ResponseWriter, tmpl *template.Template, data []client.ProductResponse) {
	w.WriteHeader(http.StatusOK)

	err := tmpl.Execute(w, data)
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

			renderNotFoundResponse(w, notFound)
			return
		}

		p, err := s.GetMapItem(q)
		if err != nil {
			c := s.GetClient()
			resp, err := c.FetchFromCMS(q)
			if err != nil {
				renderNotFoundResponse(w, notFound)
				return
			}

			fmt.Println("aa")
			s.SetMapItem(q, resp)

			renderProducts(w, show, resp)
			return
		}

		if p.IsExpired() {
			// Re-fetch from CMS.

			c := s.GetClient()
			resp, err := c.FetchFromCMS(q)
			if err != nil {
				renderNotFoundResponse(w, notFound)
				return
			}

			s.SetMapItem(q, resp)

			renderProducts(w, show, resp)
			return
		} else {
			// Render products.

			renderProducts(w, show, p.Products)
			return
		}
	}
}
