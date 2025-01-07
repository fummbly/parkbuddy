package main

import (
	"encoding/xml"
	"fmt"
	"log"

	xmlHelper "github.com/fummbly/parkbuddy/internal/xml"
)

func main() {

	data, err := xmlHelper.ReadFile()

	if err != nil {
		log.Fatalf("Failed to read file: %v\n", err)
		return
	}

	v := xmlHelper.OSM{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		log.Fatalf("Failed to unmarshal xml: %v\n", err)
		return
	}

	for _, node := range v.Nodes {
		fmt.Printf("ID: %d\n", node.ID)
		fmt.Printf("Tags: %v\n", node.Tags)

	}

}
