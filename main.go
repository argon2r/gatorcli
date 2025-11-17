package gatorcli

import (
	"fmt"
	"log"

	"github.com/argon2r/gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("kunai")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config user: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
