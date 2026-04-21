package httpserver

import (
	"errors"
	"fmt"
	"net/http"
)

func StartHttpServer() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("обработка запроса /ping")
		w.Write([]byte("hello form docker\n"))
	})

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}
