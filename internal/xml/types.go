package xml

import (
	"encoding/xml"
	"time"
)

type OSM struct {
	XMLName   xml.Name  `xml:"osm"`
	Version   string    `xml:"version,attr"`
	Generator string    `xml:"generator,attr"`
	Timestamp time.Time `xml:"timestamp,attr"`
	Bounds    Bounds    `xml:"bounds"`
	Nodes     []Node    `xml:"node"`
}

type Bounds struct {
	MinLat float32 `xml:"minlat,attr"`
	MinLon float32 `xml:"minlon,attr"`
	MaxLat float32 `xml:"maxlat,attr"`
	MaxLon float32 `xml:"maxlon,attr"`
}

type Node struct {
	ID      int     `xml:"id,attr"`
	Lat     float32 `xml:"lat,attr"`
	Lon     float32 `xml:"lon,attr"`
	Version string  `xml:"version,attr"`
	Tags    []Tag   `xml:"tag"`
}

type Tag struct {
	Key   string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}
