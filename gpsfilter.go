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
	fmt.Println("Enter points file path to process....")
	var filepath string
	filepath = "data/points.csv"
	fmt.Scan(&filepath)
	if filepath == "" {
		fmt.Println("Invalid file path")
		os.Exit(3)
	}
	csvFile, _ := os.Open(filepath)
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
	
	for idx,val:= range points{
		if idx > 0 {
			//calculate distance between two points
			dist := filter_utility.Distance(points[idx-1].Lattitude,points[idx-1].Longitude,val.Lattitude,val.Longitude)
			//time diffrence in seconds
			tf := val.Timestamp -points[idx-1].Timestamp
			points[idx].Speed = int64(math.Round(dist))/ tf
		} else {
			//Set speed 1 m/sec for the starting point
			points[idx].Speed =1
		}
		
	}

	var nonZeroPoints []point.Point
	//Filter points with zero speed
	for _,val:= range points{
		if val.Speed > 0 {
			nonZeroPoints = append(nonZeroPoints,val)
		}
	}

	fmt.Println(len(nonZeroPoints))

	//Get Standard daviation for speed
	sd:= int64(filter_utility.CalculateSDForSpeed(nonZeroPoints))
	var nonOutlierPoints []point.Point
	//Filter outlier point > 2*sd
	//2 is standard number and can be changed based on noisy data
	for _,val:= range nonZeroPoints{
		if val.Speed < 2*sd {
			nonOutlierPoints = append(nonOutlierPoints,val)
		}
	}

	//Generate the output file
	filter_utility.WritePointsResult(nonOutlierPoints)
}