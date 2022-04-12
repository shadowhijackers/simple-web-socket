package app

import (
	"log"

	"github.com/google/uuid"
)

func handleError(err error) {
	if err == nil {
		log.Fatal(err.Error())
	}
}

func generateUId() string {
	id := uuid.New()
	return id.String()
}
