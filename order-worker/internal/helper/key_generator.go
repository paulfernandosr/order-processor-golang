package helper

import "fmt"

func GenerateOrderLockKey(id string) string {
	return generateKey("lock", "order", id)
}

func GenerateOrderErrorKey(id string) string {
	return generateKey("error", "order", id)
}

func generateKey(prefix string, context string, id string) string {
	return fmt.Sprintf("%s:%s:%s", prefix, context, id)
}
