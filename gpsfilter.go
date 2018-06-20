package main

import(
    "fmt"
	"os"
	"math"
	"go-gps-filter/filter_utility"
	"go-gps-filter/point"
)



func main(){
	fmt.Println("Enter points file path to process....")
	var filepath string
	//filepath = "data/points.csv"
	fmt.Scan(&filepath)
	if filepath == "" {
		fmt.Println("Invalid file path")
		os.Exit(3)
	}
	
	//read Points file 
	points:= filter_utility.ReadPointsFile(filepath)

	//Calculate speed using distance and time 
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

	fmt.Println("Total Available  Points", len(points))

	var nonZeroPoints []point.Point
	//Filter points with zero speed
	for _,val:= range points{
		if val.Speed > 0 {
			nonZeroPoints = append(nonZeroPoints,val)
		}
	}

	fmt.Println("Total Non Zero Points:", len(nonZeroPoints))

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

	fmt.Println("Total Non outlier Points:", len(nonOutlierPoints))

	//Generate the output file
	filter_utility.WritePointsResult(nonOutlierPoints)
}