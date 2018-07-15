package data

import (
	"encoding/hex"

	"github.com/satori/go.uuid"
)

func generateID() string {
	uuid := uuid.Must(uuid.NewV4())
	return hex.EncodeToString(uuid.Bytes())
}
