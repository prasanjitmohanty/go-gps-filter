package main

import(
	"bufio"
    "encoding/csv"
    //"encoding/json"
    "fmt"
    "io"
    "log"
	"os"
	"strconv"
	"math"
	//"time"
	"go-gps-filter/filter_utility"
	"go-gps-filter/point"
)



func main(){
	csvFile, _ := os.Open("data/points.csv")
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
	
	for i,val:= range points{
		if i > 0 {
			//calculate distance between two points
			dist := filter_utility.Distance(points[i-1].Lattitude,points[i-1].Longitude,val.Lattitude,val.Longitude)
			//time diffrence in seconds
			tf := val.Timestamp -points[i-1].Timestamp
			val.Speed = int64(math.Round(dist))/ tf
		} else {
			//Set speed 1 m/sec for the starting point
			val.Speed =1
		}
		fmt.Println(val.Speed)
		//fmt.Print(i)
	}

	//Filter points with zero speed
	
}