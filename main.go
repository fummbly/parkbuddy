package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("./data/map/planet_-71.277,42.263_-70.932,42.441.osm")
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}

	fmt.Println(string(data))

}
