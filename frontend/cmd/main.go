package main

import (
	"log"
	"os"

	_ "example.com/weather"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	port, _ := os.LookupEnv("PORT")
	if port == "" {
		port = "8090"
	}
	if err := funcframework.StartHostPort("", port); err != nil {
		log.Fatalf("funcframework.StartHostPort: %v\n", err)
	}
}
