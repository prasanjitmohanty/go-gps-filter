package filter_utility

import (
    "os"
    "log"
	"encoding/csv"
	"go-gps-filter/point"
	"strconv"
	"bufio"
	"io"
)
func ReadPointsFile(filePath string) []point.Point {
	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var points []point.Point
	for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
		}

		lat,_ := strconv.ParseFloat(line[0],64)
		long,_ := strconv.ParseFloat(line[1],64)
		ts,_ := strconv.ParseInt(line[2],10,64)
		
        points = append(points, point.Point{
            Lattitude: lat,
            Longitude:  long,
			Timestamp:  ts,
		})
	}

	return points
}

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