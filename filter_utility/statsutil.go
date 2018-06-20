package filter_utility

import(
	"math"
	"go-gps-filter/point"
)

//Calculate Standard Deviation for speed
func CalculateSDForSpeed(points []point.Point) float64 {
	var sum,mean,sd float64
	for _,v := range points{
		sum = sum+float64(v.Speed);
	}
	
	mean = sum/float64(len(points))
	for _,v := range points{
		
		sd += math.Pow(float64(v.Speed) - mean, 2)
   }

   sd = math.Sqrt(sd/float64(len(points)))
   return sd
}