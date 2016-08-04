package main

import (
	"fmt"
	"log"
	"os"

	"github.com/smartystreets/smartystreets-go-sdk/us-street-api"
	"github.com/smartystreets/smartystreets-go-sdk/wireup"
)

func main() {
	log.SetFlags(log.Ltime)

	client := wireup.NewClientBuilder().
		WithSecretKeyCredential(os.Getenv("SMARTY_AUTH_ID"), os.Getenv("SMARTY_AUTH_TOKEN")).
		BuildUSStreetAPIClient()

	if err := client.Ping(); err != nil {
		fmt.Println("Ping failed:", err)
		os.Exit(1)
	} else {
		fmt.Println("Ping successful; service is reachable and responding.")
	}

	batch := us_street.NewBatch()
	for batch.Append(&us_street.Lookup{Street: "3214 N University ave", LastLine: "Provo UT 84604"}) {
		fmt.Print(".")
	}
	fmt.Println("\nBatch full, preparing to send inputs:", batch.Length())

	if err := client.SendBatch(batch); err != nil {
		log.Fatal("Error sending batch:", err)
	}

	for i, input := range batch.Records() {
		fmt.Println("Input:", i)
		for j, candidate := range input.Results {
			fmt.Println("Candidate:", j)
			fmt.Println(candidate.DeliveryLine1)
			fmt.Println(candidate.LastLine)
			fmt.Println()
		}
	}
}