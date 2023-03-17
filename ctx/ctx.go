package ctx

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			log.Println("context has been cancelled and it was not able to finish the fetch from the store")
			return
		}

		fmt.Fprint(w, data)
	}
}
