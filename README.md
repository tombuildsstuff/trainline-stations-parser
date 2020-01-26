# Trainline Stations Parser

This package allows you to parse [the Trainline.eu Stations List](https://github.com/trainline-eu/stations) into a Go object.

## Example Usage

This assumes the file `stations.csv` ([from the Trainline.eu repository](https://github.com/trainline-eu/stations)) is present within the same folder.

```go
package main

import (
    "log"
    "os"
    
    "github.com/tombuildsstuff/trainline-stations-parser/parser"
)

func main() {
    fileName := "stations.csv"
    stations, err := parser.Parse(fileName)
	if err != nil {
		log.Printf("Error parsing: %+v", err)
		os.Exit(1)
	}
	
    for _, station := range *stations {
        log.Printf("Station: %q in %q (%.6f, %.6f)", station.Name, station.Country, station.Latitude, station.Longitude)
    }
}
```