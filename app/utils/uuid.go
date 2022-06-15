package utils

import "github.com/google/uuid"

func UUIDV1() string {
	return uuid.Must(uuid.NewUUID()).String()
}

//UUIDV4 func() (string, error)
func UUIDV4() string {
	return uuid.NewString()
}
