//go:generate templ generate
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	homePage := index()
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v [%v] [%v] - GET /", r.RemoteAddr, r.UserAgent(), r.Header.Get("X-Forwarded-For"))
		homePage.Render(r.Context(), w)
	})
	//templ.Handler(homePage))
	http.HandleFunc("POST /next/{days}", func(w http.ResponseWriter, r *http.Request) {
		days, err := strconv.Atoi(r.PathValue("days"))
		if err != nil {
			http.Error(w, "Invalid days", http.StatusBadRequest)
			return
		}
		log.Printf("%v [%v] [%v] - POST /next/%v", r.RemoteAddr, r.UserAgent(), r.Header.Get("X-Forwarded-For"), days)
		later := time.Now().Add(time.Duration(days) * 24 * time.Hour)
		fmt.Fprintf(w, "In %d days it will be %s", days, later.Format("2006-01-02"))
	})

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		i := 0
		for range ticker.C {
			log.Printf("Tick %d", i)
			i++
		}
	}()

	http.ListenAndServe(":8080", nil)
}
