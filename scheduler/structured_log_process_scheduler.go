package scheduler

import (
	"log"
	"log-processor/repository"
	"time"
)

func StartLogCompletionScheduler() {
	for {
		log.Println("Running Trace Completion Check...")

		err := repository.MarkCompletedTraces()
		if err != nil {
			log.Printf("Error marking completed traces: %v", err)
		}

		log.Println("Trace Completion Check Done!")
		time.Sleep(10 * time.Second)
	}
}
