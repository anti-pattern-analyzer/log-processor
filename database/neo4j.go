package database

import (
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var Neo4jDriver neo4j.Driver

func ConnectNeo4j() {
	neo4jURI := getEnv("NEO4J_URI", "bolt://localhost:7687")
	neo4jUsername := getEnv("NEO4J_USERNAME", "neo4j")
	neo4jPassword := getEnv("NEO4J_PASSWORD", "H7t-YFBxjSBWxIFnhO7OaGAjKwk6c-q5wxq_yP95E1M")

	var err error
	Neo4jDriver, err = neo4j.NewDriver(neo4jURI, neo4j.BasicAuth(neo4jUsername, neo4jPassword, ""))
	if err != nil {
		log.Fatalf("Error connecting to Neo4j: %v", err)
	}

	err = Neo4jDriver.VerifyConnectivity()
	if err != nil {
		log.Fatalf("Error verifying Neo4j connection: %v", err)
	}

	log.Println("Connected to Neo4j at", neo4jURI)
}

// Helper function to get environment variables with a default fallback
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
