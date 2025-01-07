package xml

import (
	"fmt"
	"os"
)

func ReadFile() ([]byte, error) {
	fmt.Println("Reading file....")
	data, err := os.ReadFile("./data/map/planet_-71.277,42.263_-70.932,42.441.osm")
	if err != nil {
		return nil, err
	}

	return data, nil

}
