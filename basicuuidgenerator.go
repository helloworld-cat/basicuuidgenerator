package basicuuidgenerator

import (
	"crypto/rand"
	"fmt"
)

func NewV4() (string, error) {
	buffer := make([]byte, 16)
	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	uuid := fmt.Sprintf(
		"%X-%X-%X-%X-%X",
		buffer[0:4],
		buffer[4:6],
		buffer[6:8],
		buffer[8:10],
		buffer[10:],
	)

	return uuid, nil
}
