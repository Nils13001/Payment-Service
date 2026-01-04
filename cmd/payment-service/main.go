package main

import (
    "log"
    "net/http"
	myhttp "github.com/Nils13001/Payment-Service/internal/http"
    "github.com/go-chi/chi/v5"
)

func main() {
    // Create router (Chi)
    r := chi.NewRouter()

    // Routes will go here
    r.Get("/api/payment/health", myhttp.HealthHandler)
	r.Post("/api/payment/authorize", myhttp.AuthorizeHandler)

	//Defining starting port
    addr := ":8091"
    log.Printf("payment-service listening on %s", addr)
	
	// 		Start the server with the router as handler
    //    Using log.Fatal ensures we see startup errors clearly
    if err := http.ListenAndServe(addr, r); err != nil {
        log.Fatal(err)
    //Trial of Branch Protection
    }

}
