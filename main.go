package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Enviroment variable 'PORT' not be set")
	}
	serve := http.Server{
		Handler:      m,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Printf("server started on %v\n", port)
	err := serve.ListenAndServe()
	log.Fatal(err)

}

func handlePage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`
	w.Write([]byte(page))
}
