package filter_utility

import (
    "os"
    "log"
	"encoding/csv"
	"go-gps-filter/point"
	"strconv"
)

func WritePointsResult(points []point.Point) {
    file, err := os.Create("result.csv")
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, val := range points {
        err := writer.Write([]string{strconv.FormatFloat(val.Lattitude, 'f', 5, 64), strconv.FormatFloat(val.Longitude, 'f', 5, 64), strconv.FormatInt(val.Timestamp, 10)})
        checkError("Cannot write to file", err)
    }
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}