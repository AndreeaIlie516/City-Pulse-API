package utils

import (
	"math/rand"
	"strconv"
)

var MinRange = 10000
var MaxRange = 100000

func generateRandomID(minRange int, maxRange int) string {
	return strconv.Itoa(rand.Intn(maxRange) + minRange)
}

func isUniqueID(newID string, existingIDs []string) bool {
	for _, id := range existingIDs {
		if id == newID {
			return false
		}
	}
	return true
}

func CreateUniqueID(minRange int, maxRange int, existingIDs []string) string {
	for {
		newID := generateRandomID(minRange, maxRange)
		if isUniqueID(newID, existingIDs) {
			return newID
		}
	}
}
