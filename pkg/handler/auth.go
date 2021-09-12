package handler

import (
	"os"
)

func isUser(username string) bool {
	snId := os.Getenv("SNEJ_USERNAME")
	admId := os.Getenv("ADMIN_USERNAME")

	if username == admId || username == snId {
		return true
	}

	return false
}
