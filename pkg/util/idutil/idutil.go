package idutil

import (
	guuid "github.com/google/uuid"
)

const (
	BASIC_UUID_SIZE = 6
)

// Create UUID by prefix string and extension map
func CreateUUID(prefix string, ext map[string]interface{}) string {
	uuid := guuid.NewString()
	return prefix + "-" + uuid
}
