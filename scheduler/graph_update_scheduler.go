package scheduler

import (
	"log"
	"log-processor/services"
	"time"
)

func StartGraphUpdateScheduler() {
	for {
		log.Println("Running Graph Update for Completed Traces...")
		err := services.UpdateGraphForCompletedTraces()
		if err != nil {
			log.Printf("Error updating graph: %v", err)
		}
		log.Println("Graph Update Completed!")
		time.Sleep(30 * time.Second)
	}
}
