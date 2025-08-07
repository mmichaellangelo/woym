package main

import (
	"context"
	"fmt"
	"net/http"
)

const PORT = ":8080"

func main() {
	db, err := InitDBPool(context.Background())
	if err != nil {
		panic("Couldn't create DB Pool")
	}

	mux := http.NewServeMux()
	thoughtHandler := newThoughtHandler(db)
	mux.Handle("/thoughts", thoughtHandler)

	fmt.Println("Listening on " + PORT)
	http.ListenAndServe(PORT, mux)
}
