package hello

import (
	"fmt"
	"os"
)

func GetHello(s string) string {
	return fmt.Sprintf("Hello World from %s PORT %v", s, os.Getenv("PORT"))
}
