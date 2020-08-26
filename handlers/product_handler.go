package handlers

import (
	"fmt"
	"github.com/yigitsadic/qrmenum_client/store"
	"net/http"
)

func ProductHandler(s store.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Get("yigit")
		fmt.Fprint(w, "Hello World")
	}
}
