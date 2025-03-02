package repository

import (
	"log-processor/database"
	"log-processor/models"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func UpdateGraphForTrace(structuredLog models.StructuredLog, version string) error {
	session := database.Neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	// If the destination is "null" and it's the last log of the trace, do not insert it into Neo4j
	if structuredLog.Destination == "null" {
		isLastLog, err := IsLastLogInTrace(structuredLog.TraceID, structuredLog.SpanID)
		if err != nil {
			return err
		}

		if isLastLog {
			return nil
		}
	}

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		// Create or update service nodes
		_, err := tx.Run(`
			MERGE (s:Service {name: $source})
			MERGE (d:Service {name: $destination})
			MERGE (s)-[r:CALLS]->(d)
			ON CREATE SET r.calls = 1
			ON MATCH SET r.calls = r.calls + 1
		`, map[string]interface{}{
			"source":      structuredLog.Source,
			"destination": structuredLog.Destination,
			"version":     version,
		})
		if err != nil {
			return nil, err
		}

		// Create or update method nodes
		_, err = tx.Run(`
			MERGE (m:Method {name: $method, type: $type})
			ON CREATE SET m.calls = 0, m.total_duration = 0, m.avg_duration = 0
		`, map[string]interface{}{
			"method": structuredLog.Method,
			"type":   structuredLog.Type,
		})
		if err != nil {
			return nil, err
		}

		// Link each method to its service
		_, err = tx.Run(`
			MATCH (s:Service {name: $source})
			MATCH (m:Method {name: $method})
			MERGE (s)-[r:INVOKES]->(m)
			ON CREATE SET r.calls = 1
			ON MATCH SET r.calls = r.calls + 1
		`, map[string]interface{}{
			"source": structuredLog.Source,
			"method": structuredLog.Method,
		})
		if err != nil {
			return nil, err
		}

		// Link method-to-method calls across services
		_, err = tx.Run(`
			MATCH (m1:Method {name: $method})
			MATCH (m2:Method {name: $destination})
			MERGE (m1)-[r:CALLS]->(m2)
			ON CREATE SET r.calls = 1, r.total_duration = $duration, r.avg_duration = $duration
			ON MATCH SET r.calls = r.calls + 1, r.total_duration = r.total_duration + $duration, r.avg_duration = r.total_duration / r.calls
		`, map[string]interface{}{
			"method":      structuredLog.Method,
			"destination": structuredLog.Destination,
			"duration":    structuredLog.DurationMs,
		})

		return nil, err
	})

	return err
}
