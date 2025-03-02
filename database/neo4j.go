package database

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var Neo4jDriver neo4j.Driver

func ConnectNeo4j() {
	neo4jURI := "neo4j+s://1843ef16.databases.neo4j.io"
	neo4jUsername := "neo4j"
	neo4jPassword := "H7t-YFBxjSBWxIFnhO7OaGAjKwk6c-q5wxq_yP95E1M"

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
