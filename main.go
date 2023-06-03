package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/axiomhq/axiom-go/axiom"
	"github.com/axiomhq/axiom-go/axiom/ingest"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	axiom_api_token := os.Getenv("AXIOM_API_TOKEN")
	axiom_org_id := os.Getenv("AXIOM_ORG_ID")

	client, err := axiom.NewClient(
		axiom.SetPersonalTokenConfig(axiom_api_token, axiom_org_id),
	)

	ctx := context.Background()

	if false {
		if _, err = client.IngestEvents(ctx, "brijesh-blog-backend", []axiom.Event{
			{ingest.TimestampField: time.Now(), "foo": "bar"},
			{ingest.TimestampField: time.Now(), "bar": "foo"},
		}); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("hi")
}
